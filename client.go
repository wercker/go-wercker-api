package wercker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jtacoma/uritemplates"
)

// NewClient creates a new Client. It merges the default Config together with
// config.
func NewClient(config *Config) *Client {
	c := &Client{config: defaultConfig.Merge(config)}

	return c
}

// Client is the wercker api client.
type Client struct {
	config *Config
}

// Do makes a request to the wercker api servers.
func (c *Client) Do(method string, urlTemplate *uritemplates.UriTemplate, urlModel interface{}, payload interface{}, result interface{}) error {
	m, ok := struct2map(urlModel)
	if !ok {
		return errors.New("Invalid URL model")
	}

	var payloadReader io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		payloadReader = bytes.NewReader(b)
	}

	path, err := urlTemplate.Expand(m)
	if err != nil {
		return err
	}

	body, err := c.makeRequest(method, path, payloadReader)
	if err != nil {
		return err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, result)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) generateURL(path string) string {
	endpoint := strings.TrimRight(c.config.Endpoint, "/")
	return endpoint + path
}

// MakeRequest makes a request to the wercker API, and returns the returned
// payload
func (c *Client) makeRequest(method string, path string, payload io.Reader) ([]byte, error) {
	url := c.generateURL(path)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	if c.config.Credentials != nil {
		// Add credentials
		creds, err := c.config.Credentials.GetCredentials()
		if err != nil {
			return nil, err
		}

		if creds.Token != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", creds.Token))
		}

		if creds.Username != "" && creds.Password != "" {
			req.SetBasicAuth(creds.Username, creds.Password)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		if resp.ContentLength != 0 {
			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			if err != nil {
				return nil, err
			}

			return body, nil
		}
		return nil, nil
	}

	return nil, c.handleError(resp)
}

// ErrResponse is a generic error object using wercker api conventions.
type ErrResponse struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"error"`
	Message       string `json:"message"`
}

// Error returns the wercker error message
func (e *ErrResponse) Error() string {
	return e.Message
}

func (c *Client) handleError(resp *http.Response) error {
	if resp.ContentLength > 0 {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		// Continue if we were able to read the response
		if err == nil {
			v := &ErrResponse{}
			err := json.Unmarshal(body, v)

			// Continue if we were able to unmarshal the JSON
			if err == nil {
				return v
			}
		}
	}

	return fmt.Errorf("Unable to parse error response (status code: %d)", resp.StatusCode)
}
