package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	server := app.Listen(":3000")
	log.Fatal(server)
}
