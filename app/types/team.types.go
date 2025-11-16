package types

import (
	// "github.com/lib/pq"
	"github.com/el-hombre0/pr-reviewer-assignment/app/dal"

)

type TeamResponse struct {
	TeamName 	string 				`json:"team_name"`
	Members 	[]dal.TeamMember	`json:"members"`
	// Members []dal.TeamMember `gorm:"foreignKey:TeamID" json:"members"`
	// Members 	pq.StringArray	`gorm:"type:text[]" json:"members"`
}

type CreateTeamDTO struct {
	TeamName 	string 				`json:"team_name"`
	Members 	[]dal.TeamMember	`json:"members"`
		// Members []dal.TeamMember `gorm:"foreignKey:TeamID" json:"members"`
	// Members 	pq.StringArray	`gorm:"type:text[]" json:"members"`
}

type CreateTeamResponse struct {
	Team *TeamResponse `json:""`
}