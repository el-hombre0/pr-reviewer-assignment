package services

import (
	"github.com/el-hombre0/pr-reviewer-assignment/app/dal"
	"github.com/el-hombre0/pr-reviewer-assignment/app/types"
	"github.com/el-hombre0/pr-reviewer-assignment/utils"
	"github.com/gofiber/fiber/v2"
)

func SetUserIsActive(c *fiber.Ctx) error {
	b := new(types.SetUserActiveDTO)

	if err := utils.ParseBody(c, b); err != nil {
		return err
	}
	userID := b.UserID

	if userID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid user id")
	}


	d := &dal.User{}

	if err := dal.FindUserByID(d, userID).Error; err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "User not found")
	}

	newUser := &dal.User{
		UserID: d.UserID,
		Username: d.Username,
		TeamName: d.TeamName,
		IsActive: b.IsActive,
	}

	err := dal.UpdateUser(userID, newUser).Error

	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}
	
	return c.JSON(&types.UserResponse{
		UserID: newUser.UserID,
		Username: newUser.Username,
		TeamName: newUser.TeamName,
		IsActive: newUser.IsActive,
	})

}

func GetUserReview(c *fiber.Ctx) error {
	return nil
}

func CreateUser(c *fiber.Ctx) error{
	b := new(types.CreateUserDTO)

	err := utils.ParseBody(c, b)
	if err != nil {
		return err
	}

	d := &dal.User{
		UserID: b.UserID,
		Username: b.Username,
		TeamName: b.TeamName,
		IsActive: b.IsActive,
	}

	if err := dal.CreateUser(d).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(
		&types.UserCreateResponse{
			User: &types.UserResponse{
				UserID: b.UserID,
				Username: b.Username,
				TeamName: b.TeamName,
				IsActive: b.IsActive,
			},
		},
	)
}