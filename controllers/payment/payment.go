package payment

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/haris/dto"
	"github.com/haris/models"
)

func CreatePayment(c *gin.Context)  {
	var input models.InputPayment;
	var customer models.Customer;
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "format not valid",
		})
		return
	}

	//get customer id by token
	token := c.GetHeader("X-API-TOKEN");
	if err := models.DB.First(&customer, "token = ?", token).Error ; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return
	}

	//check token expired at
	status:= customer.TokenExpiredAt.After(time.Now())
	if status == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "token expired",
		})
		return
	} 


	payload := models.Payment{Amount: input.Amount, CustomerID: customer.Username, MerchantId: input.MerchantId, TransactionId: input.TransactionId, Status: input.Status}

	if err := models.DB.Create(&payload).Error ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	response := dto.PaymentResponse{
		ID: payload.ID,
		Amount: payload.Amount,
		CustomerID: payload.CustomerID,
		MerchantId: payload.MerchantId,
		TransactionId: payload.TransactionId,
		Status: payload.Status,
		Updated: payload.Updated,
		Created: payload.Created,
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "succes create payment",
		"data" : response,
	})

}
	func FindAllPayment(c *gin.Context)  {
	var payment []models.Payment
	var customer models.Customer
	var merchant models.Merchant
	token := c.GetHeader("X-API-TOKEN")
	


	errC := models.DB.First(&customer, "token = ?", token).Error ;
	errM := models.DB.First(&merchant, "token = ?", token).Error;
	if errC != nil && errM != nil  {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return
	}

	if errC == nil {
	//check token expired at
	statusC:= customer.TokenExpiredAt.After(time.Now())
	if statusC == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "token expired",
		})
		return
	} 
	}

	if errM == nil {
	statusM:= merchant.TokenExpiredAt.After(time.Now())
	if statusM == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "token expired",
		})
		return
	} 
	}
	
	if err := models.DB.Find(&payment).Where("customer_id = ?", customer.Username ).Or("merchant_id = ?", merchant.Username).Error ; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "payment not found",
		})
		return
	}

	//untk map to response
	var responses []dto.PaymentResponse
	for _, v := range payment {
		response := dto.PaymentResponse{
			ID: v.ID,
			Amount: v.Amount,
			CustomerID: v.CustomerID,
			MerchantId: v.MerchantId,
			TransactionId: v.TransactionId,
			Status: v.Status,
			Updated: v.Updated,
			Created: v.Created,
		}
		responses = append(responses, response)
	}


	c.JSON(http.StatusOK, gin.H{
		"message" : "OK",
		"data" : responses,
	})
}

func FindPaymentStatus(c *gin.Context)  {
	var response dto.PaymentResponse
	var payment models.Payment
	var customer models.Customer
	var merchant models.Merchant
	token := c.GetHeader("X-API-TOKEN")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return
	}


	errC := models.DB.First(&customer, "token = ?", token).Error;
	errM := models.DB.First(&merchant, "token = ?", token).Error;

	if errC != nil && errM != nil  {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "unauthorized",
		})
		return
	}

	if errC == nil {
	statusC:= customer.TokenExpiredAt.After(time.Now())
	if statusC == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "token expired",
		})
		return
	} 
	}

	if errM == nil {
	statusM:= merchant.TokenExpiredAt.After(time.Now())
	if statusM == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "token expired",
		})
		return
	} 
	}

	//get payment by id
	paymentId := c.Param("id")
	if err := models.DB.First(&payment, paymentId).Error ; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	response = dto.PaymentResponse{
		ID: payment.ID,
		Amount: payment.Amount,
		CustomerID: payment.CustomerID,
		MerchantId: payment.MerchantId,
		TransactionId: payment.TransactionId,
		Status: payment.Status,
		Created: payment.Created,
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "OK",
		"data" : response,
	})
	
}