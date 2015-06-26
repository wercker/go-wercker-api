package wercker

import "github.com/jtacoma/uritemplates"

func NewClient(options Options) *Client {

	return &Client{}
}

type Client struct {
	options Options
}

func (c *Client) makeRequest(method string, url uritemplates.UriTemplate) {

}
