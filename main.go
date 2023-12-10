package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haris/controllers/auth"
	"github.com/haris/models"
)
  
  func main() {

	r := gin.Default()
	models.ConnectDatabase()


	r.POST("/api/auth/register", auth.RegisterCustomer)
	r.POST("/api/auth/login", auth.Login)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
  }

