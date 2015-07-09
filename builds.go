package wercker

type buildsService struct {
	client *Client
}

type GetBuildOptions struct {
	Config *Config `map:"-"`

	BuildID string `map:"buildId"`
}

func (c *buildsService) Get(options *GetBuildOptions) (*Build, error) {
	method := "GET"
	template := routes["buildService.GetBuild"]

	var payload interface{} = nil
	var headers map[string]string = nil
	result := &Build{}

	err := c.client.makeRequest(method, template, options, payload, headers, options.Config, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type FetchForApplicationOptions struct {
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

func (c *buildsService) FetchForApplication(options *FetchForApplicationOptions) ([]*BuildSummary, error) {
	method := "GET"
	template := routes["applications.FetchApplicationBuilds"]

	var payload interface{} = nil
	var headers map[string]string = nil
	result := []*BuildSummary{}

	err := c.client.makeRequest(method, template, options, payload, headers, options.Config, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
