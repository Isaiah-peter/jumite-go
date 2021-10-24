package main

import (
	"awesomeProject/jumite/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("welcome")
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	routes.UserRoute(app)
	routes.AuthRoute(app)
	routes.OrderRoute(app)
	routes.MessageRoute(app)
	routes.ProductRoute(app)
	routes.CommentRoute(app)

	app.Listen(":6000")
}
