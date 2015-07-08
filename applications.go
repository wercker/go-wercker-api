package wercker

type GetApplicationOptions struct {
	Config *Config

	// Required
	Owner string
	Name  string
}

func (c *Client) GetApplication(options *GetApplicationOptions) (*Application, error) {
	method := "GET"
	template := routes["applications.GetApplication"]

	var payload interface{} = nil
	var headers map[string]string = nil
	result := &Application{}

	err := c.makeRequest(method, template, options, payload, headers, options.Config, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type FetchApplicationBuildsOptions struct {
	Config *Config `map:"-"`

	// Required
	Owner string `map:"owner"`
	Name  string `map:"name"`

	// Optional
	Branch string `map:"branch,omitempty"`
	Commit string `map:"commit,omitempty"`
	Limit  string `map:"limit,omitempty"`
	Result string `map:"result,omitempty"`
	Skip   string `map:"skip,omitempty"`
	Sort   string `map:"sort,omitempty"`
	Stack  string `map:"stack,omitempty"`
	Status string `map:"status,omitempty"`
}

func (c *Client) FetchApplicationBuilds(options *FetchApplicationBuildsOptions) ([]*BuildSummary, error) {
	method := "GET"
	template := routes["applications.FetchApplicationBuilds"]

	var payload interface{} = nil
	var headers map[string]string = nil
	result := []*BuildSummary{}

	err := c.makeRequest(method, template, options, payload, headers, options.Config, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
