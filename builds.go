package wercker

type GetBuildOptions struct {
	*Options
	BuildID string
}

func (c *Client) GetBuild(options *GetBuildOptions) (*Build, error) {
	method := "GET"
	urlTemplate := routes["Build.GetBuild"]
	urlModel := make(map[string]interface{})

	var payload interface{} = nil
	var headers map[string]string = nil
	result := &Build{}

	err := c.makeRequest(method, urlTemplate, urlModel, payload, headers, options.Options, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) TriggerBuild() (*Build, error) {
	return nil, nil
}
