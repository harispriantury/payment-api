package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  

var DB *gorm.DB
func ConnectDatabase()  {
	dsn := "host=localhost user=postgres password=password dbname=go-procurement port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if (err != nil) {
		log.Fatal("Failed to connect database postgresql")
	}

	db.AutoMigrate(Customer{})
	DB = db
	fmt.Println("Success connect database")
}