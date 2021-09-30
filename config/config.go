package config

import (
	"fmt"
	"go-ecommerce/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		GetEnvVariable("DB_USERNAME"),
		GetEnvVariable("DB_PASSWORD"),
		GetEnvVariable("DB_HOST"),
		GetEnvVariable("DB_PORT"),
		GetEnvVariable("DB_NAME"),
	)
	var e error
	DB, e = gorm.Open(mysql.Open(connDB), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(
		&models.Provinces{},
		&models.City{},
		&models.Users{},
		&models.Categories{},
		&models.Products{},
		&models.PaymentMethods{},
		&models.Transactions{},
		&models.TransactionDetails{},
		&models.Carts{})
}
