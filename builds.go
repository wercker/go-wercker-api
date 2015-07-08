package wercker

type GetBuildOptions struct {
	*Options
	BuildID string
}

func (c *Client) GetBuild(options *GetBuildOptions) (*Build, error) {
	method := "GET"
	template := routes["builds.GetBuild"]

	var payload interface{} = nil
	var headers map[string]string = nil
	result := &Build{}

	err := c.makeRequest(method, template, options, payload, headers, options.Options, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
