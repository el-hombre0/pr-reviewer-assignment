package main

import (
	"log"
	"os"

	"github.com/el-hombre0/pr-reviewer-assignment/app/dal"
	"github.com/el-hombre0/pr-reviewer-assignment/app/routes"
	"github.com/el-hombre0/pr-reviewer-assignment/config/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// type Repository struct {
// 	DB *gorm.DB
// }


// func (r *Repository) CreatePR(context *fiber.Ctx) error{
// 	pr := models.PullRequest{}

// 	err := context.BodyParser(&pr)

// 	if err != nil{
// 		log.Fatal("En error occured while parsing pull request")
// 		context.Status(http.StatusUnprocessableEntity).JSON(
// 			&fiber.Map{"message": "En error occured while processing request!"},
// 		)
// 		return err
// 	}

// 	err = r.DB.Create(&pr).Error

// 	if err != nil {
// 		context.Status(http.StatusBadRequest).JSON(
// 			&fiber.Map{"message": "En error occured while creating pull request!"},)
// 	}

// 	context.Status(http.StatusOK).JSON(
// 			&fiber.Map{"message": "Pull request created successfully!"},)
// 	return nil
// }

// func (r *Repository)CreateTeam(context *fiber.Ctx) error{
// 	team := models.Team{}

// 	err := context.BodyParser(&team)

// 	if err != nil{
// 		log.Print("En error occured while parsing team")
// 		context.Status(http.StatusUnprocessableEntity).JSON(
// 			&fiber.Map{"message": "En error occured while processing request!"},
// 		)
// 		return err
// 	}

// 	err = r.DB.Create(&team).Error

// 	if err != nil {
// 		context.Status(http.StatusBadRequest).JSON(
// 			&fiber.Map{"message": "En error occured while creating team!"},)
// 	}


// 	context.Status(http.StatusOK).JSON(
// 		&fiber.Map{"team": team},
// 	)
// 	return nil

// }

// func (r *Repository) GetTeamByTeamName(context *fiber.Ctx) error{
// 	teamName := context.Params("team_name")
// 	teamModel := &models.Team{}
// 	if teamName == ""{
// 		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "TeamName can not be empty!", 
// 		})
// 		return nil
// 	}

// 	fmt.Println("The TeamName is", teamName)

// 	err := r.DB.Where("team_name = ?", teamName).First(teamModel).Error

// 	if err != nil {
// 		context.Status(http.StatusBadRequest).JSON(
// 			&fiber.Map{"message": "En error occured while geting a team!"},)
// 			return err
// 	}

// 	context.Status(http.StatusOK).JSON(
// 		&fiber.Map{"team_name": teamName, "members": teamModel},
// 	)
// 	return nil
// }


func main(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}

	config := &database.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_DBNAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	err = database.NewConnection(config)

	if err != nil {
		log.Fatal("Error while loading the database!")
	}

	err = database.Migrate(
		&dal.PullRequest{},
		&dal.TeamMember{},
		&dal.Team{},
		&dal.Users{},
	)

	if err != nil {
		log.Fatal("An error occured while migrating pull requests!")
	}

	app := fiber.New()

	// TODO
	// app := fiber.New(fiber.Config{
	// 	ErrorHandler: utils.ErrorHandler,
	// })

	routes.PullRequestRoutes(app)
	routes.TeamRoutes(app)
	routes.UsersRoutes(app)
	app.Listen(":8080")
}