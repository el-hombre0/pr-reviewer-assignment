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

	if err := dal.CreatePullRequest(d).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
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
	b := new(types.PullRequstMergeDTO)

	if err := utils.ParseBody(c, b); err != nil {
		return err
	}
	pullRequestID := b.PullRequestID

	if pullRequestID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid pull request id")
	}


	d := &dal.PullRequest{}

	if err := dal.FindPullRequestByID(d, pullRequestID).Error; err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Pull request not found")
	}

	newPR := &dal.PullRequest{
		PullRequestID: d.PullRequestID,
		PullRequestName: d.PullRequestName,
		AuthorID: d.AuthorID,
		Status: dal.PullRequestStatusMerged,
		AssignedReviewers: d.AssignedReviewers,
		CreatedAt: d.CreatedAt,
		MergedAt: time.Now(),
	}

	err := dal.UpdatePullRequest(pullRequestID, newPR).Error

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}
	
	return c.JSON(&types.PullRequestMergeResponse{
		PullRequest: &types.PullRequestResponse{
				PullRequestID: newPR.PullRequestID,
				PullRequestName: newPR.PullRequestName, 
				AuthorID: newPR.AuthorID, 
				Status: string(newPR.Status), 
				AssignedReviewers: newPR.AssignedReviewers,
			},
	})
}

func ReassignPullRequest(c *fiber.Ctx) error{
	return nil
}