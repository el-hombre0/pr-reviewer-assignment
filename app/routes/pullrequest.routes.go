package routes

import (
	"github.com/el-hombre0/pr-reviewer-assignment/app/services"
	"github.com/gofiber/fiber/v2"
)

func PullRequestRoutes(app fiber.Router){
	prApi := app.Group("/pullRequest")
	// prApi := app.Group("/pullRequest").Use(middlware.Validation)
	prApi.Post("/create", services.CreatePullRequest)
	prApi.Post("/merge", services.MergePullRequest)
	prApi.Post("/reassign", services.ReassignPullRequest)
}