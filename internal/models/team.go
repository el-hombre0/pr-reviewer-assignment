package models

import (
	"gorm.io/gorm"
)

type Team struct {
	TeamName string 	`json:"team_name"`
	Members []TeamMember `gorm:"type:text" json:"members"`
}

func MigrateTeam(db *gorm.DB) error{
	err := db.AutoMigrate(&Team{})
	return err
}