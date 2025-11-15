package routes

import (
	"github.com/el-hombre0/pr-reviewer-assignment/app/services"
	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app fiber.Router){
	usersApi := app.Group("/users")
	usersApi.Post("/setIsActive", services.SetUserIsActive)
	usersApi.Get("/getReview", services.GetUserReview)
}