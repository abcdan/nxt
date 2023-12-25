package routes

import (
	"nxt/helper"

	"github.com/gofiber/fiber/v2"
)

func TotalLinks(c *fiber.Ctx) error {
	links, err := helper.TotalLinks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot fetch total links",
		})
	}
	return c.JSON(fiber.Map{
		"links": links,
	})
}

func StatisticsRoutes(app *fiber.App) {
	app.Get("/api/statistics/links", TotalLinks)
}

