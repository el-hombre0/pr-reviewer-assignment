package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type status string

const (
	StateOpen status = iota
	StateMerged
)

var statusName = map[status]string{
	StateOpen: "OPEN",
	StateMerged: "MERGED",
}

type PullRequest struct {
	pull_request_id 	string	`json: "pull_request_id"`
	pull_request_name 	string	`json: "pull_request_name"`
	author_id 			string	`json: "author_id"`
	status 				string	`json: "status"`
}

func (r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/pullRequest/create", r.CreatePR)
	api.Post("/pullRequest/merge", r.MergePR)
	api.Post("/pullRequest/reassign", r.ReassignPR)
}

func main(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Error while loading the database!")
	}

	var r Repository = Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}