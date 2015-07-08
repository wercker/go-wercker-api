package wercker

type GetApplicationOptions struct {
	*Options
	Owner string
	Name  string
}

func (c *Client) GetApplication(options *GetApplicationOptions) (*Application, error) {
	method := "GET"
	template := routes["applications.GetApplication"]

	var payload interface{} = nil
	var headers map[string]string = nil
	result := &Application{}

	err := c.makeRequest(method, template, options, payload, headers, options.Options, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type GetApplicationBuildsOptions struct {
	*Options

	// Required
	Owner string
	Name  string

	// Optional
	Branch string `uri:"branch,omitempty"`
	Commit string `uri:"commit,omitempty"`
	Limit  string `uri:"limit,omitempty"`
	Result string `uri:"result,omitempty"`
	Skip   string `uri:"skip,omitempty"`
	Sort   string `uri:"sort,omitempty"`
	Stack  string `uri:"stack,omitempty"`
	Status string `uri:"status,omitempty"`
}

func (c *Client) GetApplicationBuilds(options *GetApplicationBuildsOptions) ([]*BuildSummary, error) {
	method := "GET"
	template := routes["applications.GetApplicationBuilds"]

	var payload interface{} = nil
	var headers map[string]string = nil
	result := []*BuildSummary{}

	err := c.makeRequest(method, template, options, payload, headers, options.Options, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
