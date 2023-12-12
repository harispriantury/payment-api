package models

type Payment struct {
	ID uint `gorm:"primaryKey"`
	Amount int `gorm:"not null"`
	CustomerID string `gorm:"not null"`
	MerchantId string `gorm:"not null"`
	Customer Customer `gorm:"references:Username;not null"`
	Merchant Merchant `gorm:"references:Username;not null"`
	TransactionId string `gorm:"not null"`
	Status Status `gorm:"not null"`
	Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
	Created   int64 `gorm:"autoCreateTime"`
	
  }

  type InputPayment struct {
	Amount int `gorm:"not null"`
	MerchantId string `gorm:"not null"`
	TransactionId string `gorm:"not null"`
	Status Status `gorm:"not null"`
  }



  type Status string

  const (
	  Pending Status = "pending"
	  Approved Status = "failed"
	  Rejected Status = "success"
  )