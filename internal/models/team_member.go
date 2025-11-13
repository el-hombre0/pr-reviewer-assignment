package models

import (
	"gorm.io/gorm"
)

type TeamMember struct {
	UserId string 	`json:"user_id"`
	Username string `json:"username"`
	IsActive bool 	`json:"is_active"`
}

func MigrateTeamMember(db *gorm.DB) error{
	err := db.AutoMigrate(&TeamMember{})
	return err
}
