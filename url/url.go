package url

import (
	"github.com/daniferdinandall/ws-dhs-p3/controller"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/swagger"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Homepage2)
	// page.Get("/docs/*", swagger.HandlerDefault)
}
