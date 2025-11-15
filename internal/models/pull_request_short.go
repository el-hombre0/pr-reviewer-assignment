package models

type PullRequestShortStatus string

const (
	PullRequestShortStatusOpen PullRequestShortStatus = "OPEN"
	PullRequestShortStatusMerged PullRequestShortStatus = "OPEN"

)

type PullRequestShort struct {
	PullRequestID string `json:"pull_request_id"`
    PullRequestName string `json:"pull_request_name"`
    AuthorID string `json:"author_id"`
    Status PullRequestShortStatus `json:"status"`
}