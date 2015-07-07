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

func NewClient(options *Options) *Client {
	return &Client{options: options}
}

type Client struct {
	options *Options
}

type ErrResponse struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"error"`
	Message       string `json:"message"`
}

func (e *ErrResponse) Error() string {
	return e.Message
}

func (c *Client) generateUrl(path string) string {
	endpoint := strings.TrimRight(c.options.Endpoint, "/")
	return endpoint + path
}

func (c *Client) MakeRequest(method string, path string) ([]byte, error) {
	client := &http.Client{}

	url := c.generateUrl(path)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if c.options.Creds != nil {
		// Add credentials
		creds, err := c.options.Creds.GetCredentials()
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

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Going to call %d", resp.StatusCode)

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {

		if resp.ContentLength > 0 {
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

func (c *Client) makeRequest(method string, template *uritemplates.UriTemplate, urlModel map[string]interface{}, payload interface{}, headers map[string]string, options *Options, result interface{}) error {
	path, err := template.Expand(urlModel)
	if err != nil {
		return err
	}

	fmt.Printf("Going to call %s", path)

	body, err := c.MakeRequest(method, path)

	// Continue if we were able to read the response
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)

	// Return an error if we were unable to unmarshal
	if err != nil {
		return err
	}
	return nil
}
