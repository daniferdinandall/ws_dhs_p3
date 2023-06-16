package controller

import (
	// "github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
)

func Homepage2(c *fiber.Ctx) error {
	// ipaddr := musik.GetIPaddress()
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Hello World",
	})
}
