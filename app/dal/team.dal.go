package dal

import (
	"github.com/el-hombre0/pr-reviewer-assignment/config/database"
	"gorm.io/gorm"
)

type Team struct {
	TeamName string 	`json:"team_name"`
	Members []TeamMember `gorm:"type:text" json:"members"`
}

func CreateTeam(team *Team) *gorm.DB {
	return database.DB.Create(team)
}