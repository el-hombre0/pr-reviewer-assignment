// Для тестирования

package types

type UserResponse struct {
	UserID string		`json:"user_id"`
	Username string		`json:"username"`
	TeamName string 	`json:"team_name"`
	IsActive bool		`json:"is_active"`
}

type CreateUserDTO struct {
	UserID string		`json:"user_id"`
	Username string		`json:"username"`
	TeamName string 	`json:"team_name"`
	IsActive bool		`json:"is_active"`
}

type UserCreateResponse struct {
	User *UserResponse `json:"user"`
}