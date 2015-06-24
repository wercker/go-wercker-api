package wercker

func NewClient(options Options) *Client {

	return &Client{}
}

type Client struct {
	options Options
}
