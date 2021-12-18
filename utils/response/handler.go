package response

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(fiber.Map{
		"Status":  "Success",
		"Message": message,
		"Data":    data,
	})
}
