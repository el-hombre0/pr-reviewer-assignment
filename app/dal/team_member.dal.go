package dal

import (
	"github.com/el-hombre0/pr-reviewer-assignment/config/database"
	"gorm.io/gorm"
)

type TeamMember struct {
	ID 			uint	`gorm:"primaryKey" json:"id"`
	UserId 		string 	`json:"user_id"`
	Username 	string	`json:"username"`
	IsActive 	bool 	`json:"is_active"`
	TeamID 		string	`json:"team_id"`
}

func CreateTeamMember(teamMember *TeamMember) *gorm.DB {
	return database.DB.Create(teamMember)
}