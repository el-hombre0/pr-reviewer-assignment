package dal

import (
	"github.com/el-hombre0/pr-reviewer-assignment/config/database"
	"gorm.io/gorm"
)

type Team struct {
	ID			uint			`gorm:"primaryKey" json:"id"`
	TeamName 	string 			`json:"team_name"`
	// Members []TeamMember `gorm:"type:text" json:"members"`
	Members 	[]TeamMember 	`gorm:"foreignKey:TeamID" json:"members"`
}

func CreateTeam(team *Team) *gorm.DB {
	return database.DB.Create(team)
}

func FindTeam(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Team{}).Take(dest, conds...)
}

func FindTeamByName(dest interface{}, teamName interface{}) *gorm.DB {
	return FindTeam(dest, "team_name = ?", teamName)
}