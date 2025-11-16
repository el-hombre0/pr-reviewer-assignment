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

// Поиск PullRequest, соответствующего заданным условиям
func FindPullRequest(dest interface{}, conds ...interface{}) *gorm.DB{
	return database.DB.Model(&PullRequest{}).Take(dest, conds)
}

// func FindPullRequestByUser(dest interface{}, prIden interface{}, userIden interface{}) *gorm.DB{
// 	return FindPullRequest(dest, "pull_request_id = ? and ")
// }

func FindPullRequestByID(dest interface{}, prIden interface{}) *gorm.DB {
	return FindPullRequest(dest, "pull_request_id = ?", prIden)
}

func UpdatePullRequest(prIden interface{}, data interface{}) *gorm.DB {
	return database.DB.Model(&User{}).Where("pull_request_id = ?", prIden).Updates(data)
}
