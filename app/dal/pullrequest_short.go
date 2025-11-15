package dal

type PullRequestShortStatus string

const (
	PullRequestShortStatusOpen PullRequestShortStatus = "OPEN"
	PullRequestShortStatusMerged PullRequestShortStatus = "OPEN"

)

type PullRequestShort struct {
	PullRequestID string 			`json:"pull_request_id"`
    PullRequestName string 			`json:"pull_request_name"`
    AuthorID string 				`json:"author_id"`
    Status PullRequestShortStatus 	`json:"status"`
}


// func CreatePullRequestShort(pr *PullRequestShort) *gorm.DB{
// 	return database.DB.Create(pr)
// }