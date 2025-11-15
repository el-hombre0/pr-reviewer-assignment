package dal

import (
	"time"

	"github.com/el-hombre0/pr-reviewer-assignment/config/database"
	"gorm.io/gorm"
	"github.com/lib/pq"
)



type PullRequestStatus string

const (
	PullRequestStatusOpen   PullRequestStatus = "OPEN"
	PullRequestStatusMerged	PullRequestStatus = "MERGED"
)

type PullRequest struct {
	PullRequestID     string              	`json:"pull_request_id"`
	PullRequestName   string              	`json:"pull_request_name"`
	AuthorID          string              	`json:"author_id"`
	Status            PullRequestStatus   	`json:"status"`
	AssignedReviewers pq.StringArray        `gorm:"type:text[]" json:"assigned_reviewers"`
	CreatedAt         time.Time          	`json:"createdAt,omitempty"`
	MergedAt          time.Time          	`json:"mergedAt,omitempty"`
}

func CreatePullRequest(pr *PullRequest) *gorm.DB{
	return database.DB.Create(pr)
}


// func MergePullRequest(pr *PullRequest) *gorm.DB{
// 	return database.DB.Merge(pr)
// }

// func ReassignPullRequest(pr *PullRequest) *gorm.DB{
// 	return database.DB.Reassign(pr)
// }