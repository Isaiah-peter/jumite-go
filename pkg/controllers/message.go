package controllers

import (
	"awesomeProject/jumite/pkg/models"
	"awesomeProject/jumite/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CreateMessage(c *fiber.Ctx) error {
	token := utils.UseToken(c)
	message := new(models.Message)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	if err := c.BodyParser(message); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	message.SenderId = verifiedID
	db.Save(&message)
	return c.JSON(&message)
}

func GetMessage(c *fiber.Ctx) error {
	var message []models.Order
	token := utils.UseToken(c)
	ids := c.Params("id")
	id, err := strconv.ParseInt(ids, 0, 0)
	if err != nil {
		panic(err)
	}
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	if token["IsAdmin"] == true || verifiedID == id  {
		db.Where("order_id=?", id).Find(&message)
		return c.JSON(&message)
	}

	return c.SendString("you are not an admin")
}
