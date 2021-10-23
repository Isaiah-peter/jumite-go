package controllers

import (
	"awesomeProject/jumite/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.CustomerProduct)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Save(&product)
	return c.JSON(&product)
}

func GetProduct(c *fiber.Ctx) error {
	var product []models.CustomerProduct
	db.Find(&product)
	return c.JSON(&product)
}
