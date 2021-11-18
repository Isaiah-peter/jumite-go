package controllers

import (
	"awesomeProject/jumite/pkg/models"
	"awesomeProject/jumite/pkg/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
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

func CreateConversation(c *fiber.Ctx) error {
	token := utils.UseToken(c)
	conversation := new(models.Conversation)
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}

	if err := c.BodyParser(conversation); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	conversation.UserId = verifiedID
	db.Save(&conversation)
	return c.JSON(&conversation)
}

func GetConversation(c *fiber.Ctx) error {
	var conversation []models.Conversation
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

	if token["IsAdmin"] == true || verifiedID == id {
		db.Where("reciever_id=?", id).Or("sender_id=?", id).Find(conversation)
		return c.JSON(&conversation)
	}

	return c.SendString("you are not an admin")
}

func GetMessage(c *fiber.Ctx) error {
	var message []models.Message
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
	if token["IsAdmin"] == true || verifiedID == id {
		db.Where("conversation_id=?", id).Find(&message)
		return c.JSON(&message)
	}

	return c.SendString("you are not an admin")
}
