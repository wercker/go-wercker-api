package wercker

type GetBuildOptions struct {
	Config  *Config
	BuildID string
}

func (c *Client) GetBuild(options *GetBuildOptions) (*Build, error) {
	method := "GET"
	template := routes["builds.GetBuild"]

	var payload interface{} = nil
	var headers map[string]string = nil
	result := &Build{}

	err := c.makeRequest(method, template, options, payload, headers, options.Config, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
