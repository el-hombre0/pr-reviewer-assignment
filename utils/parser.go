package utils

import "github.com/gofiber/fiber/v2"

// Обертка, чтобы не писать много if condition
func ParseBody(c *fiber.Ctx, body interface{}) *fiber.Error {
	err := c.BodyParser(body);
	if err != nil {
		return fiber.ErrBadRequest
	}
	return nil
}

// func ParseBodyAndValidate(c *fiber.Ctx, body interface{}) *fiber.Error {
// 	err := ParseBody(c, body)
// 	if err != nil {
// 		return err
// 	}
// 	return Validate(body)
// }