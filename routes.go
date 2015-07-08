package wercker

import "github.com/jtacoma/uritemplates"

// routes is a map containing all UriTemplates indexed by name.
var routes map[string]*uritemplates.UriTemplate = make(map[string]*uritemplates.UriTemplate)

func init() {
	// Add templates to the route map
	addURITemplate("applications.GetApplication", "/api/v3/applications{/owner,name}")
	addURITemplate("applications.FetchApplicationBuilds", "/api/v3/applications{/owner,name}/builds{?commit,branch,status,limit,skip,sort,result}")
	addURITemplate("builds.GetBuild", "/api/v3/builds{/buildId}")
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
