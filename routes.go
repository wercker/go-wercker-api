package wercker

import "github.com/jtacoma/uritemplates"

// routes is a map containing all UriTemplates indexed by name.
var routes map[string]*uritemplates.UriTemplate = make(map[string]*uritemplates.UriTemplate)

func init() {
	// Add templates to the route map
	addURITemplate("applications.GetApplication", "/api/v3/applications{/Owner,Name}")
	addURITemplate("applications.GetApplicationBuilds", "/api/v3/applications{/Owner,Name}/builds{?commit,branch,status,limit,skip,sort,result}")
	addURITemplate("builds.GetBuild", "/api/v3/builds{/BuildID}")
}

// addURITemplate adds rawTemplate to routes using name as the key (Should only
// be used from init()).
func addURITemplate(name, rawTemplate string) {
	uriTemplate, err := uritemplates.Parse(rawTemplate)
	if err != nil {
		panic(err)
	}
	routes[name] = uriTemplate
}
