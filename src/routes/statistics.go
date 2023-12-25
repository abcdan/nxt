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

func LinkClicks(c *fiber.Ctx) error {
	shortcode := c.Params("shortcode")
	link, err := helper.GetLinkByShortcode(shortcode)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot fetch link",
		})
	}
	clicks := helper.Clicks(link)
	return c.JSON(fiber.Map{
		"clicks": clicks,
	})
}

func StatisticsRoutes(app *fiber.App) {
	app.Get("/api/statistics/links", TotalLinks)
	app.Get("/api/statistics/links/:shortcode", LinkClicks)
}

