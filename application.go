package wercker

import "github.com/jtacoma/uritemplates"

// applicationTemplates contains all UriTemplates indexed by name.
var applicationTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(applicationTemplates, "GetApplication", "/api/v3/applications{/owner,name}")
}

// GetApplicationOptions are the options associated with Client.GetApplication
type GetApplicationOptions struct {
	// Required
	Owner string `map:"owner"`
	Name  string `map:"name"`
}

// GetApplication will retrieve a single Application
func (c *Client) GetApplication(options *GetApplicationOptions) (*Application, error) {
	method := "GET"
	template := applicationTemplates["GetApplication"]

	result := &Application{}
	err := c.Do(method, template, options, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetApplicationOptions are the options associated with Client.GetApplication
type GetApplicationsOptions struct {
	Limit string `map:"limit,omitempty"`
	Skip  int    `map:"skip,omitempty"`
	Sort  int    `map:"sort,omitempty"`
	Stack string `map:"stack,omitempty"`
}

// GetApplication will retrieve a single Application
func (c *Client) GetApplications(options *GetApplicationsOptions) ([]Application, error) {
	method := "GET"
	template := applicationTemplates["GetApplications"]

	result := []Application{}
	err := c.Do(method, template, options, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
