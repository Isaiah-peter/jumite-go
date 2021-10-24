package controllers

import (
	"awesomeProject/jumite/pkg/models"
	"awesomeProject/jumite/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Save(&comment)
	return c.JSON(&comment)
}

func GetComment(c *fiber.Ctx) error {
	var comment []models.Comment
	utils.UseToken(c)
	db.Order("created_at DESC").Limit(3).Find(&comment)
	return c.JSON(&comment)
}
