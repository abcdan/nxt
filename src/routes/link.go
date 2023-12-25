package routes

import (
	"nxt/helper"
	"nxt/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LinkRoutes(app *fiber.App) {
	app.Post("/api/link", func(c *fiber.Ctx) error {
		link := new(models.Link)
		if err := c.BodyParser(link); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		link.ShortCode = helper.GenerateShortCode(6)
		link.CreatedAt = time.Now()
		link.IP = helper.HashIP(c.IP())

		if link.PassCode != nil {
			hashedPassCode, err := helper.EncryptPassCode(*link.PassCode)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Cannot encrypt password",
				})
			}
			link.PassCode = &hashedPassCode
		}

		return c.JSON(fiber.Map{
			"short_code": link.ShortCode,
			"redirects_to": link.URL,
			"url": "https://" + os.Getenv("DOMAIN") + "/" + link.ShortCode,
			"created_at": link.CreatedAt,
		})
	})

	app.Delete("/api/link/:shortCode", func(c *fiber.Ctx) error {
		shortCode := c.Params("shortCode")
		link := new(models.Link)
		link, err := helper.GetLinkByShortcode(shortCode)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Link not found",
			})
		}
		if link != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Link not found",
			})
		}

		if link.IP == helper.HashIP(c.IP()) || (link.PassCode != nil && helper.CheckPassCode(c.Get("passcode"), *link.PassCode)) {
			err = helper.DeleteLinkByShortcode(shortCode)
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Link not found",
				})
			}
		} else {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Link not found",
			})
		}

		err = helper.DeleteLinkByShortcode(shortCode)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Link not found",
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	})
}
