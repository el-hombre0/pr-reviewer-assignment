package services

import (
	"github.com/el-hombre0/pr-reviewer-assignment/app/dal"
	"github.com/el-hombre0/pr-reviewer-assignment/app/types"
	"github.com/el-hombre0/pr-reviewer-assignment/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateTeam(c *fiber.Ctx) error {
	b := new(types.CreateTeamDTO)

	err := utils.ParseBody(c, b)
	if err != nil {
		return err
	}

	d := &dal.Team{
		TeamName: b.TeamName,
		Members: b.Members,
	}

	if err := dal.CreateTeam(d).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.TeamResponse{
				TeamName: d.TeamName,
				Members: d.Members,
		},
	)
}

func GetTeamByName(c *fiber.Ctx) error {
	teamName := c.Params("teamName")

	if teamName == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid teamName")
	}

	d := &types.TeamResponse{}

	err := dal.FindTeamByName(d, teamName).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Comand not found")
	}

	return c.JSON(&types.TeamResponse{
		TeamName: d.TeamName,
		Members: d.Members,
	})
}