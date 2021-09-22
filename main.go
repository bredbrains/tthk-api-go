package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"tthkAPI/views"
)

func main() {
	app := fiber.New()
	app.Get("/changes", views.Changes)

	server := app.Listen(":3000")
	log.Fatal(server)
}
