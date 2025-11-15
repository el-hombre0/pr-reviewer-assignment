package dal

import (
	"github.com/el-hombre0/pr-reviewer-assignment/config/database"
	"gorm.io/gorm"
)

type User struct {
	UserID string		`json:"user_id"`
	Username string		`json:"username"`
	TeamName string 	`json:"team_name"`
	IsActive bool		`json:"is_active"`
}

func CreateUser(team *User) *gorm.DB {
	return database.DB.Create(team)
}