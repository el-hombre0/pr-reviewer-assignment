package types

type PullRequestResponse struct {
	PullRequestID string `json:"pull_request_id"`
    PullRequestName string `json:"pull_request_name"`
    AuthorID string `json:"author_id"`
    Status string `json:"status"`
    AssignedReviewers []string `json:"assigned_reviewers"`
}

type CreateDTO struct {
	PullRequestID string `json:"pull_request_id"`
    PullRequestName string `json:"pull_request_name"`
	AuthorID string `json:"author_id"` 
	// TODO
    // AuthorID string `json:"author_id" validate:"required,min=3,max=150"` 
}

type PullRequestCreateResponse struct {
	PullRequest *PullRequestResponse `json:"pr"`
}

type PullRequstMergeDTO struct {
	PullRequestID string `json:"pull_request_id"`
}

type PullRequestMergeResponse struct {
	PullRequest *PullRequestResponse `json:"pr"`
}

type PullRequestReassignDTO struct {

}

type PullRequestReassignResponse struct {
	
}