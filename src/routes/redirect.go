package routes

import (
	"nxt/helper"

	"github.com/gofiber/fiber/v2"
)

func RedirectRoutes(app *fiber.App) {
	app.Get("/:shortCode", func(c *fiber.Ctx) error {
	
		shortCode := c.Params("shortCode")
		
		preClaimed := map[string]string{
			"source":     "https://github.com/abcdan/nxt",
			"sourcecode": "https://github.com/abcdan/nxt",
			"code":       "https://github.com/abcdan/nxt",
			"github":     "https://github.com/abcdan/nxt",
			"src":        "https://github.com/abcdan/nxt",
			"deploy":     "https://railway.app/template/gSvwgO?referralCode=QsZ-bg",
			"railway":    "https://railway.app/template/gSvwgO?referralCode=QsZ-bg",
			"host":       "https://railway.app/template/gSvwgO?referralCode=QsZ-bg",
			"up":         "https://railway.app/template/gSvwgO?referralCode=QsZ-bg",
		}

		if url, ok := preClaimed[shortCode]; ok {
			return c.Redirect(url, fiber.StatusFound)
		}

		link, err := helper.GetLinkByShortcode(shortCode)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
		}
		helper.ClickToLink(link)
		return c.Redirect(link.URL, fiber.StatusFound)
	})
}


