package controllers

import (
	"awesomeProject/jumite/pkg/models"
	"awesomeProject/jumite/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var (
	NewOrder models.Order
)

func CreateOrder(c *fiber.Ctx) error {
	token := utils.UseToken(c)
	order := new(models.Order)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	if err := c.BodyParser(order); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	order.UserId = verifiedID
	db.Save(&order)
	return c.JSON(&order)
}

func GetOrder(c *fiber.Ctx) error {
	var order []models.Order
	token := utils.UseToken(c)

	if token["IsAdmin"] == true {
		db.Where("delete=?", "").Preload("Messages").Find(&order)
		return c.JSON(&order)
	}

	return c.SendString("you are not an admin")
}

func GetOrderById(c *fiber.Ctx) error {
	token:= utils.UseToken(c)
	ids := c.Params("id")
	id, err := strconv.ParseInt(ids, 0, 0)
	if err != nil {
		panic(err)
	}
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	if NewOrder.Delete != "" {
		return c.Status(500).SendString("you have delete this order")
	}
	if verifiedID == id || token["IsAdmin"] == true {
		db.Where("user_id=?", id).Preload("Messages").Find(&NewOrder)
		return c.JSON(&NewOrder)
	}

	return c.Status(fiber.StatusUnauthorized).SendString("this is not your order")
}

func UpdateOrder(c *fiber.Ctx) error {
	order := &models.Order{}
	 utils.UseToken(c)
	id := c.Params("id")
	db.First(&order, id)
	if order.Delete != "" {
		return c.Status(500).SendString("you have delete this order")
	}
	if err := c.BodyParser(order); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Save(&order)
	return c.JSON(&order)
}

func DeleteOrder(c *fiber.Ctx) error {
	order := &models.Order{}
	utils.UseToken(c)
	id := c.Params("id")
	db.First(&order, id)
	if err := c.BodyParser(order); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	order.Delete = "deleted"
	db.Save(&order)
	return c.JSON(&order)
}

