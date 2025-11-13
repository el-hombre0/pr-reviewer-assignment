package internal.models
/*
User:
      type: object
      required: [ user_id, username, team_name, is_active ]
      properties:
        user_id:
          type: string
        username:
          type: string
        team_name:
          type: string
        is_active:
          type: boolean
*/
import (
	"gorm.io/gorm"
)

type Users struct {
	UserID uint			`gorm:"primary key;autoIncrement" json:"user_id"`
	Username string		`json:"username"`
	TeamName string 	`json:"team_name"`
	IsActive boolean	`json:"is_active"`
}

func MigrateUsers(db *grom.DB) error{
	err := db.AutoMigrate(&Users)
	return err
}