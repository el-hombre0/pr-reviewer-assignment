package routes

import (
	"github.com/el-hombre0/pr-reviewer-assignment/app/services"
	"github.com/gofiber/fiber/v2"
)

func TeamRoutes(app fiber.Router){
	teamApi := app.Group("/team")
	teamApi.Post("/add", services.CreateTeam)
	teamApi.Get("/get", services.GetTeamByName)
}