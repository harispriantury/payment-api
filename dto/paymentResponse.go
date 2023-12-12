package dto

import (
	"github.com/haris/models"
)

type PaymentResponse struct {
	ID uint `gorm:"primaryKey"`
	Amount int `gorm:"not null"`
	CustomerID string `gorm:"not null"`
	MerchantId string `gorm:"not null"`
	TransactionId string `gorm:"not null"`
	Status models.Status `gorm:"not null"`
	Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
	Created   int64 `gorm:"autoCreateTime"`
  }