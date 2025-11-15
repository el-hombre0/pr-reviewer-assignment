package services

import (
	"time"

	"github.com/el-hombre0/pr-reviewer-assignment/app/dal"
	"github.com/el-hombre0/pr-reviewer-assignment/app/types"
	"github.com/el-hombre0/pr-reviewer-assignment/utils"
	"github.com/gofiber/fiber/v2"
)

func CreatePullRequest(c *fiber.Ctx) error{
	b := new(types.CreateDTO)

	err := utils.ParseBody(c, b)
	if err != nil {
		return err
	}
	
	d := &dal.PullRequest{
		PullRequestID: b.PullRequestID,
		PullRequestName: b.PullRequestName,
		AuthorID: b.AuthorID,
		Status: dal.PullRequestStatusOpen,
		AssignedReviewers: AssignReviewers(),
		CreatedAt: time.Now(),
		MergedAt: time.Time{},
	}

	return c.JSON(
		&types.PullRequestCreateResponse{
			PullRequest: &types.PullRequestResponse{
				PullRequestID: d.PullRequestID,
				PullRequestName: d.PullRequestName, 
				AuthorID: d.AuthorID, 
				Status: string(d.Status), 
				AssignedReviewers: d.AssignedReviewers,
			},
		},
	)
}

func AssignReviewers() []string{
	// TODO
	assignedReviewers := []string {
		"1",
		"2",
	}
	return assignedReviewers
}

func MergePullRequest(c *fiber.Ctx) error{
	return nil
}

func ReassignPullRequest(c *fiber.Ctx) error{
	return nil
}