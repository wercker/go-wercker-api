package wercker

type buildsService struct {
	client *Client
}

// GetBuildOptions are the options associated with buildsService.Get
type GetBuildOptions struct {
	Config *Config `map:"-"`

	BuildID string `map:"buildId"`
}

// Get will retrieve a single Build
func (c *buildsService) Get(options *GetBuildOptions) (*Build, error) {
	method := "GET"
	template := routes["buildService.GetBuild"]

	result := &Build{}
	err := c.client.makeRequest(method, template, options, nil, nil, options.Config, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FetchForApplicationOptions are the options associated with
// buildsService.FetchForApplication
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

// FetchForApplication Fetch all builds for a certain application
func (c *buildsService) FetchForApplication(options *FetchForApplicationOptions) ([]*BuildSummary, error) {
	method := "GET"
	template := routes["applications.FetchApplicationBuilds"]

	result := []*BuildSummary{}
	err := c.client.makeRequest(method, template, options, nil, nil, options.Config, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
