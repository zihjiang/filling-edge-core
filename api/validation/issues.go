

package validation

type Issues struct {
	IssueCount     int                `json:"issueCount"`
	PipelineIssues []Issue            `json:"pipelineIssues"`
	StageIssuesMap map[string][]Issue `json:"stageIssues"`
}

func NewIssues(issues []Issue) *Issues {
	issuesObj := &Issues{
		IssueCount:     len(issues),
		PipelineIssues: make([]Issue, 0),
		StageIssuesMap: make(map[string][]Issue),
	}
	for _, issue := range issues {
		if len(issue.InstanceName) == 0 {
			issuesObj.PipelineIssues = append(issuesObj.PipelineIssues, issue)
		} else {
			if issuesObj.StageIssuesMap[issue.InstanceName] == nil {
				issuesObj.StageIssuesMap[issue.InstanceName] = make([]Issue, 0)
			}
			issuesObj.StageIssuesMap[issue.InstanceName] = append(issuesObj.StageIssuesMap[issue.InstanceName], issue)
		}
	}
	return issuesObj
}
