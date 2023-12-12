package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haris/controllers/auth"
	"github.com/haris/controllers/payment"
	"github.com/haris/models"
)
  
  func main() {

	r := gin.Default()

	models.ConnectDatabase()
	


	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server running succesfully",
		})
	})

	r.POST("/api/auth/register/customer", auth.RegisterCustomer)
	r.POST("/api/auth/register/merchant", auth.RegisterMerchant)
	r.POST("/api/auth/login", auth.Login)
	r.DELETE("/api/auth/logout", auth.Logout)
	r.POST("/api/payment", payment.CreatePayment)
	r.GET("/api/payments", payment.FindAllPayment)
	r.GET("/api/payments/:id", payment.FindPaymentStatus)

	r.Run() // listen and serve on 0.0.0.0:8080
	
  }

