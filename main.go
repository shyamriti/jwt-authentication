package main

import (
	"JWT-authentication/controller"
	database "JWT-authentication/database"
	"JWT-authentication/models"
	"log"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// api := r.Group("/api")
	// {
	// 	public := api.Group("/public")
	// 	{
	// 		public.POST("/login", controller.Login)
	// 		public.POST("/signup", controller.SignUp)
	// 	}
	// }
	r.POST("/login", controller.Login)
	r.POST("/signup", controller.SignUp)
	return r
}
func main() {
	err := database.Connection()
	if err != nil {
		log.Fatalln("could not create database", err)
	}
	database.Db.AutoMigrate(&models.User{})
	r := setupRouter()
	r.Run(":8080")
}
