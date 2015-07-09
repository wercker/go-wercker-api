package wercker

type applicationsService struct {
	client *Client
}

// GetApplicationOptions are the options associated with applicationsService.Get
type GetApplicationOptions struct {
	Config *Config `map:"-"`

	// Required
	Owner string `map:"owner"`
	Name  string `map:"name"`
}

// Get will retrieve a single Application
func (c *applicationsService) Get(options *GetApplicationOptions) (*Application, error) {
	method := "GET"
	template := routes["applications.GetApplication"]

	result := &Application{}
	err := c.client.makeRequest(method, template, options, nil, nil, options.Config, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
