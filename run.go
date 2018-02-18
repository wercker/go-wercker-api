package wercker

import "github.com/jtacoma/uritemplates"

// runTemplates contains all UriTemplates indexed by name.
var runTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(runTemplates, "GetRuns", "/api/v3/runs{?applicationId,pipelineId,limit,skip,sort,status,result,branch,commit,sourceRun,author}")
	addURITemplate(runTemplates, "GetRun", "/api/v3/runs{/runId}")
	addURITemplate(runTemplates, "GetRunSteps", "/api/v3/runs{/runId}/steps")
	addURITemplate(runTemplates, "TriggerNewRun", "/api/v3/runs")
}

// GetRunsOptions are the options associated with Client.GetRuns
type GetRunsOptions struct {
	ApplicationID string `map:"applicationId,omitempty"`
	PipelineID    string `map:"pipelineId,omitempty"`
	Limit         int    `map:"limit,omitempty"`
	Skip          int    `map:"skip,omitempty"`
	Sort          string `map:"sort,omitempty"`
	Status        string `map:"status,omitempty"`
	Result        string `map:"result,omitempty"`
	Branch        string `map:"branch,omitempty"`
	Commit        string `map:"commit,omitempty"`
	SourceRun     string `map:"sourceRun,omitempty"`
	Author        string `map:"author,omitempty"`
}

// GetRuns will retrieve a list of Runs
func (c *Client) GetRuns(options *GetRunsOptions) ([]RunSummary, error) {
	method := "GET"
	template := runTemplates["GetRuns"]

	result := []RunSummary{}
	err := c.Do(method, template, options, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetRunOptions are the options associated with Client.GetRun
type GetRunOptions struct {
	RunID string `map:"runId"`
}

// GetRun will retrieve a single of Run
func (c *Client) GetRun(options *GetRunOptions) (*Run, error) {
	method := "GET"
	template := runTemplates["GetRun"]

	result := &Run{}
	err := c.Do(method, template, options, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
