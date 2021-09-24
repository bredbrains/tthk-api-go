package main

import (
	"github.com/bredbrains/tthk-api-go/views"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	app.Get("/changes", views.Changes)

	server := app.Listen(":3000")
	log.Fatal(server)
}
