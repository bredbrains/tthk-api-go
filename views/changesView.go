package views

import (
	"github.com/gofiber/fiber/v2"
	"tthkAPI/clients"
)

func Changes(c *fiber.Ctx) error {
	client := clients.ChangesClient{}
	changes := client.Parse()
	return c.JSON(changes)
}
