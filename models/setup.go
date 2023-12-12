package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  

var DB *gorm.DB
func ConnectDatabase()  {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error when get env")
	}
	
	// dsn := "host=localhost user=postgres password=password dbname=go-procurement port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	HOST := os.Getenv("HOST");
	USER := os.Getenv("USER_DB");
	PASSWORD := os.Getenv("PASSWORD");
	DB_NAME := os.Getenv("DB_NAME");
	PORT := os.Getenv("DB_PORT");


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", HOST, USER, PASSWORD, DB_NAME, PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if (err != nil) {
		log.Fatal("Failed to connect database postgresql")
	}

	db.AutoMigrate(Customer{})
	db.AutoMigrate(Merchant{})
	db.AutoMigrate(Payment{})
	DB = db
	fmt.Println("Success connect database")
}
