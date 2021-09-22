package main

import (
	"awesomeProject/jumite/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("welcome")
	})

	routes.UserRoute(app)
	routes.AuthRoute(app)

	app.Listen(":6000")
}
