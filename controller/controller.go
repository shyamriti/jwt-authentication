package controller

import (
	"JWT-authentication/auth"
	databese "JWT-authentication/database"
	"JWT-authentication/models"
	"JWT-authentication/repo"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	err = user.HashPassword(user.Password)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	resp, err := repo.CreatedUserRecord(user)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	c.JSON(200, resp)

}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User

	err := c.ShouldBind(&payload)
	if err != nil {
		fmt.Printf("err1: %v\n", err)
	}
	result := databese.Db.Where("email=?", payload.Email).Select("*").First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}
	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	c.JSON(200, tokenResponse)

}
