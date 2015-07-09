package wercker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jtacoma/uritemplates"
)

// NewClient creates a new Client. It merges the default Config together with
// config.
func NewClient(config *Config) *Client {
	c := &Client{
		config: defaultConfig.Merge(config),
	}

	c.Applications = &applicationsService{client: c}
	c.Builds = &buildsService{client: c}
	c.Tokens = &tokensService{client: c}

	return c
}

// Client is the wercker api client.
type Client struct {
	config *Config

	Applications *applicationsService
	Builds       *buildsService
	Tokens       *tokensService
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

func generateURL(path string, config *Config) string {
	endpoint := strings.TrimRight(config.Endpoint, "/")
	return endpoint + path
}

// MakeRequest makes a request to the wercker API, and returns the returned
// payload
func (c *Client) MakeRequest(method string, path string, override *Config) ([]byte, error) {
	config := c.config.Merge(override)

	url := generateURL(path, config)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if config.Credentials != nil {
		// Add credentials
		creds, err := config.Credentials.GetCredentials()
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

	resp, err := config.HTTPClient.Do(req)
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

	return errors.New("Invalid status code received")
}

func (c *Client) makeRequest(method string, template *uritemplates.UriTemplate, urlModel interface{}, payload interface{}, headers map[string]string, override *Config, result interface{}) error {
	m, ok := struct2map(urlModel)
	if !ok {
		return errors.New("Invalid URL model")
	}

	path, err := template.Expand(m)
	if err != nil {
		return err
	}

	body, err := c.MakeRequest(method, path, override)
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
