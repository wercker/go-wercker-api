package wercker

type applicationsService struct {
	client *Client
}

type GetApplicationOptions struct {
	Config *Config `map:"-"`

	// Required
	Owner string `map:"owner"`
	Name  string `map:"name"`
}

func (c *applicationsService) Get(options *GetApplicationOptions) (*Application, error) {
	method := "GET"
	template := routes["applications.GetApplication"]

	var payload interface{} = nil
	var headers map[string]string = nil
	result := &Application{}

	err := c.client.makeRequest(method, template, options, payload, headers, options.Config, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
