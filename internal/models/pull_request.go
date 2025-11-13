package models

import (
	"time"
	"gorm.io/gorm"
)

/*
PullRequest:
      type: object
      required: [ pull_request_id, pull_request_name, author_id, status, assigned_reviewers]
      properties:
        pull_request_id:
          type: string
        pull_request_name:
          type: string
        author_id:
          type: string
        status:
          type: string
          enum: [OPEN, MERGED]
        assigned_reviewers:
          type: array
          items:
            type: string
          description: user_id назначенных ревьюверов (0..2)
        createdAt:
          type: string
          format: date-time
          nullable: true
        mergedAt:
          type: string
          format: date-time
          nullable: true
*/

// PullRequestStatus представляет возможные статусы pull request
type PullRequestStatus string

const (
	PullRequestStatusOpen   PullRequestStatus = "OPEN"
	PullRequestStatusMerged	PullRequestStatus = "MERGED"
)

// PullRequest основная модель Pull Request
type PullRequest struct {
	PullRequestID     uint              	`gorm:"primary key;autoIncrement" json:"pull_request_id"`
	PullRequestName   string              	`json:"pull_request_name"`
	AuthorID          string              	`json:"author_id"`
	Status            PullRequestStatus   	`json:"status"`
	AssignedReviewers []string            	`gorm:"type:text" json:"assigned_reviewers"`
	CreatedAt         *time.Time          	`json:"createdAt,omitempty"`
	MergedAt          *time.Time          	`json:"mergedAt,omitempty"`
}

// CreatePullRequestRequest DTO для создания Pull Request
// type CreatePullRequestRequest struct {
// 	PullRequestName   string   	`json:"pull_request_name" validate:"required,min=1,max=255"`
// 	AuthorID          string   	`json:"author_id" validate:"required,uuid"`
// 	AssignedReviewers []string `json:"assigned_reviewers" validate:"max=2"`
// }

// // UpdatePullRequestRequest DTO для обновления Pull Request
// type UpdatePullRequestRequest struct {
// 	PullRequestName   string            	`json:"pull_request_name,omitempty" validate:"omitempty,min=1,max=255"`
// 	Status            PullRequestStatus 	`json:"status,omitempty" validate:"omitempty,oneof=OPEN MERGED"`
// 	AssignedReviewers []string          	`json:"assigned_reviewers,omitempty" validate:"omitempty,max=2"`
// }

// // PullRequestResponse DTO для ответа API
// type PullRequestResponse struct {
// 	PullRequestID     string            	`json:"pull_request_id"`
// 	PullRequestName   string            	`json:"pull_request_name"`
// 	AuthorID          string            	`json:"author_id"`
// 	Status            PullRequestStatus 	`json:"status"`
// 	AssignedReviewers []string          	`json:"assigned_reviewers"`
// 	CreatedAt         *time.Time        	`json:"createdAt,omitempty"`
// 	MergedAt          *time.Time        	`json:"mergedAt,omitempty"`
// }

func MigratePullRequest(db *gorm.DB) error{
	err := db.AutoMigrate(&PullRequest{})
	return err
}