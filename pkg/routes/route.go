package routes

import (
	"awesomeProject/jumite/pkg/controllers"
	"awesomeProject/jumite/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

var AuthRoute = func(route *fiber.App) {
	route.Post("/register", controllers.CreateUser)
	route.Post("/login", controllers.Login)
}

var UserRoute = func(route *fiber.App) {
	route.Get("/user", utils.Authuser(), controllers.GetUser)
	route.Get("/user/:id", utils.Authuser(), controllers.GetUserById)
	route.Put("/user/:id", utils.Authuser(), controllers.UpdateUser)
	route.Delete("/user/:id", utils.Authuser(), controllers.DeleteUser)
}

var OrderRoute = func(route *fiber.App) {
	route.Post("/order", utils.Authuser(), controllers.CreateOrder)
	route.Get("/order", utils.Authuser(), controllers.GetOrder)
	route.Get("/order/:id", utils.Authuser(), controllers.GetOrderById)
	route.Put("/order/:id", utils.Authuser(), controllers.UpdateOrder)
	route.Patch("/order/:id", utils.Authuser(), controllers.DeleteOrder)
}

var MessageRoute = func(route *fiber.App) {
	route.Post("/message", utils.Authuser(), controllers.CreateMessage)
	route.Get("/message/:id", utils.Authuser(), controllers.GetMessage)
}


