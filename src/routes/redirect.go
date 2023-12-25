package routes

import (
	"nxt/helper"

	"github.com/gofiber/fiber/v2"
)

func RedirectRoutes(app *fiber.App) {
	app.Get("/:shortCode", func(c *fiber.Ctx) error {
		shortCode := c.Params("shortCode")
		link, err := helper.GetLinkByShortcode(shortCode)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
		}
		return c.Redirect(link.URL, fiber.StatusFound)
	})
}

