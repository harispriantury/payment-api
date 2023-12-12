package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/haris/models"
	"github.com/haris/utils"
)


func RegisterCustomer(c *gin.Context) {
	var input models.InputRegister
	var payload models.Customer

	if err := c.ShouldBindJSON(&input); err != nil {
	   c.JSON(http.StatusBadRequest, gin.H{"errors": "format not valid"})
	   return
	}

	token, _ := utils.HashPassword(input.Password)

	payload = models.Customer{
		Name: input.Name,
		Username: input.Username,
		Password: token,
	}

	 if err := models.DB.Create(&payload).Error ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	 }

	c.JSON(http.StatusOK, gin.H{
		"data" : "register succesfully",
	})
}

func RegisterMerchant(c *gin.Context) {
	var input models.InputRegister
	var payload models.Merchant
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "format not valid"})
		return
	}

	token, _ := utils.HashPassword(input.Password)
	
	payload = models.Merchant{
			Name: input.Name,
			Username: input.Username,
			Password: token,
	}
	 if err := models.DB.Create(&payload).Error ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	 }

	c.JSON(http.StatusOK, gin.H{
		"data" : "register Merchant succesfully",
	})
}

func Login(c *gin.Context) {
	var input models.InputLogin
	var customer models.Customer
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "format not correctly",
		})
		return
	}

	if err := models.DB.Where("username = ?", input.Username).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}

	//cek password
	if !utils.CheckPasswordHash(input.Password, customer.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credential",
		})
		return
	}

	token, _ := utils.HashPassword(customer.Password)


	expiredAt := time.Now().Add(60 * time.Minute)

	if err := models.DB.Model(customer).Updates(models.Customer{Token: token, TokenExpiredAt: expiredAt}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login succesfully",
		"token" : token,
	})
}

func Logout (c *gin.Context) {
	var customer models.Customer;
	token := c.GetHeader("X-API-TOKEN")


	if err := models.DB.Where("token = ?", token).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "user not found",
		})
		return
	}

	customer.Token = ""
	customer.TokenExpiredAt = time.Now()

	if err := models.DB.Save(customer).Error ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "logout succesfully",
	})
	
}

// func FindAll(c *gin.Context)  {
// 	var products []models.Product
// 	result := database.DB.Find(&products)

// 	if result.Error != nil {
// 		// Tangani kesalahan database
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	if result.RowsAffected == 0 {
// 		// Tidak ada data ditemukan
// 		c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
// 		return
// 	} 	

// 	c.JSON(http.StatusOK, gin.H{
// 		"data : ": &products,
// 		"message : " : "success get all data",
// })
// }

// func FindById(c *gin.Context)  {
// 	var product models.Product

// 	id := c.Param("id")
	

// 	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "product not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data" : &product,
// 	})
	

// }

// func CreateProduct(c *gin.Context) {
// 	var input models.Product;

// 	if err := c.ShouldBindJSON(&input).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error" : "request so bad",
// 		})
// 		return
// 	}

// 	newProduct := models.Product{Name: input.Name, Price: input.Price}
// 	database.DB.Create(&newProduct);

// 	c.JSON(http.StatusOK, gin.H{"data": &newProduct})
// }

// func DeleteProduct(c *gin.Context) {
// 	var product models.Product
	

// 	if err := database.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error" : "product not found",
// 		})
// 		return

// 	}
	
// 	database.DB.Delete(&product)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message" : "product successfully deleted",
// 	})
// }

// func UpdateProduct(c *gin.Context) {
// 	var product models.Product


// 	//get product yang mau diupdate
// 	if err := database.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "record not found",
// 		})
// 		return
// 	}

// 	//valiate input
// 	var input models.Product
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error" : "input format is bad",
// 		})
// 		return
// 	}

// 	database.DB.Model(&product).Updates(models.Product{Name: input.Name, Price: input.Price})
// 	c.JSON(http.StatusOK, gin.H{
// 		"data" : "data successfully updated",
// 	})

// }


