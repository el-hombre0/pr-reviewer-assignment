package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	storage "github.com/el-hombre0/pr-reviewer-assignment/storage/database"
	models "github.com/el-hombre0/pr-reviewer-assignment/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// type status string

// const (
// 	StateOpen status = iota
// 	StateMerged
// )

// var statusName = map[status]string{
// 	StateOpen: "OPEN",
// 	StateMerged: "MERGED",
// }

// type PullRequest struct {
// 	pull_request_id 	string	`json: "pull_request_id"`
// 	pull_request_name 	string	`json: "pull_request_name"`
// 	author_id 			string	`json: "author_id"`
// 	status 				string	`json: "status"`
// }

func (r *Repository) CreatePR(context *fiber.Ctx) error{
	pr := models.PullRequest{}

	err := context.BodyParser(&pr)

	if err != nil{
		log.Fatal("En error occured while parsing pull request")
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "En error occured while processing request!"},
		)
		return err
	}

	err = r.DB.Create(&pr).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "En error occured while creating pull request!"},)
	}

	context.Status(http.StatusOK).JSON(
			&fiber.Map{"message": "Pull request created successfully!"},)
	return nil
}

func (r *Repository)CreateTeam(context *fiber.Ctx) error{
	team := models.Team{}

	err := context.BodyParser(&team)

	if err != nil{
		log.Print("En error occured while parsing team")
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "En error occured while processing request!"},
		)
		return err
	}

	err = r.DB.Create(&team).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "En error occured while creating team!"},)
	}


	context.Status(http.StatusOK).JSON(
		&fiber.Map{"team": team},
	)
	return nil

}

func (r *Repository) GetTeamByTeamName(context *fiber.Ctx) error{
	teamName := context.Params("team_name")
	teamModel := &models.Team{}
	if teamName == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "TeamName can not be empty!", 
		})
		return nil
	}

	fmt.Println("The TeamName is", teamName)

	err := r.DB.Where("team_name = ?", teamName).First(teamModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "En error occured while geting a team!"},)
			return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"team_name": teamName, "members": teamModel},
	)
	return nil

}

func (r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api/v1")
	api.Post("/pullRequest/create", r.CreatePR)
	// api.Post("/pullRequest/merge", r.MergePR)
	// api.Post("/pullRequest/reassign", r.ReassignPR)
	api.Post("/team/add", r.CreateTeam)
	api.Get("/team/get", r.GetTeamByTeamName)
}

func main(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}

	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_DBNAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Error while loading the database!")
	}

	err = models.MigratePullRequest(db)

	if err != nil {
		log.Fatal("An error occured while migrating pull requests!")
	}

	err = models.MigrateTeam(db)

	if err != nil {
		log.Fatal("An error occured while migrating teams!")
	}

	var r Repository = Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}