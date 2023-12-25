package routes

import (
	"net/url"
	"nxt/helper"
	"nxt/models"
	"os"
	"strconv"
	"strings"
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

		u, err := url.Parse(link.URL)
		if err != nil || u.Scheme == "" || u.Host == "" || strings.Contains(link.URL, os.Getenv("DOMAIN")) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid URL or URL contains the domain",
			})
		}

		shortCodeLength, _ := strconv.Atoi(os.Getenv("SHORTCODE_LENGTH"))
		for i := 0; i < 100; i++ {
			link.ShortCode = helper.GenerateShortCode(shortCodeLength)
			if _, err := helper.GetLinkByShortcode(link.ShortCode); err != nil {
				break
			}
		}

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

		err := helper.InsertLink(link)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cannot insert link into database",
			})
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
		link, err := helper.GetLinkByShortcode(shortCode)
		if err != nil {
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
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	})
}
