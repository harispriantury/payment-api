package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haris/models"
)


func RegisterCustomer(c *gin.Context) {
	var input models.Customer

	if err := c.ShouldBindJSON(&input); err != nil {
	   c.JSON(http.StatusBadRequest, gin.H{"errors": "cannot registered customer"})
	   return
	}

	 if err := models.DB.Create(&input).Error ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	 }

	c.JSON(http.StatusOK, gin.H{
		"data" : "register succesfully",
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

	if customer.Password != input.Password {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "failed authenticated",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"data" : "login successfully",
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


