package controllers

import (
	"awesomeProject/jumite/pkg/config"
	"awesomeProject/jumite/pkg/models"
	"awesomeProject/jumite/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var (
	db      = config.GetDB()
	NewUser models.User
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	user.Password = hashPassword
	db.Create(&user)
	return c.JSON(&user)
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	u := models.FindOne(user.Email, user.Password)
	return c.JSON(u)
}

func GetUser(c *fiber.Ctx) error {
	var user []models.User
	token := utils.UseToken(c)
	if token["IsAdmin"] == true {
		db.Find(&user)
		return c.JSON(&user)
	}

	return c.SendString("you are not an admin")
}

func GetUserById(c *fiber.Ctx) error {
	utils.UseToken(c)
	id := c.Params("id")
	db.Find(&NewUser, id)
	return c.JSON(&NewUser)
}

func DeleteUser(c *fiber.Ctx) error {
	token := utils.UseToken(c)
	userid := c.Params("id")
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}
	 verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}

	if verifiedID == id {
		db.First(&NewUser, id)
		db.Delete(&NewUser)
		c.Status(200)
		return c.SendString("you have delete this account")
	}

	if token["IsAdmin"] == true {
		db.First(&NewUser, id)
		db.Delete(&NewUser)
		c.Status(200)
		return c.SendString("you have delete this account")
	}
	return c.SendString("you are not an admin and you did not own this account")
}

func UpdateUser(c *fiber.Ctx) error {
	user := &models.User{}
	token := utils.UseToken(c)
	userid := c.Params("id")
	id, err := strconv.ParseInt(userid, 0, 0)
	if err != nil {
		panic(err)
	}
	verifiedID, err := strconv.ParseInt(fmt.Sprintf("%.f", token["UserID"]), 0, 0)
	if err != nil {
		panic(err)
	}
	db.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User is not available")
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if token["IsAdmin"] == true || verifiedID == id {
		fmt.Println(user.Password)
		password, err := utils.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = password
		db.Save(&user)
		return c.JSON(&user)
	}

	return c.Status(fiber.StatusUnauthorized).SendString("you don't own this account and you ale not an admin")
}
