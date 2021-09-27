package models

import (
	"awesomeProject/jumite/pkg/config"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
)

type User struct {
	gorm.Model
	UserName    string `json:"user_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	IsAdmin     bool   `json:"is_admin" asn1:"default:false"`
	PhoneNumber string `json:"phone_number"`
	Order       []Order
}

type Token struct {
	UserID  int64
	IsAdmin bool
	jwt.StandardClaims
}

func init() {
	config.Connect()
	db = config.GetDB()
	if err := db.AutoMigrate(&User{}); err != nil {
		return
	}
}

func FindOne(email string, password string) map[string]interface{} {
	var user = &User{}
	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expireAt := time.Now().Add(time.Hour * 24).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false,"password": user.Password ,"message": "Invalid login credentials. Please try again"}
		return resp
	}

	fmt.Println(errf)

	tk := &Token{
		UserID:  int64(user.ID),
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		panic(err)
	}

	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	return resp
}
