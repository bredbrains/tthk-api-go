package views

import (
	"github.com/bredbrains/tthk-api-go/clients"
	"github.com/gofiber/fiber/v2"
)

func Changes(c *fiber.Ctx) error {
	client := clients.ChangesClient{}
	changes := client.Parse()
	return c.JSON(changes)
}
