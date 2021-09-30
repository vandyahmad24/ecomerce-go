package models

import (
	sql "database/sql"

	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name" form:"name" gorm:"size:255" `
	Email       string `json:"email" form:"email" gorm:"size:255"`
	Password    string `json:"password" form:"password" gorm:"size:255"`
	ProvincesId uint   `json:"provinces_id" form:"provinces_id" `
	Provinces   Provinces
	CityId      uint `json:"city_id" form:"city_id" `
	City        City
	ZipCode     uint           `json:"zip_code,string" form:"zip_code"`
	Address     string         `json:"address" form:"address"`
	PhoneNumber string         `json:"phone_number" gorm:"size:255,unique" form:"phone_number"`
	StoreName   sql.NullString `json:"store_name" form:"store_name" gorm:"size:255"`
	StoreStatus bool           `json:"store_status" form:"store_status" gorm:"default:0"`
	OtpCode     sql.NullString `json:"otp_code" form:"otp_code" gorm:"size:14"`
	IsActive    bool           `json:"is_active" form:"is_active" gorm:"default:0"`
	IsAdmin     bool           `json:"is_admin" form:"is_admin" gorm:"default:0"`
	Token       string         `json:"token" form:"token"`
}

type UserRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Password    string `json:"password" form:"password" validate:"required"`
	CityId      uint   `json:"city_id,string" form:"city_id" validate:"required" `
	ZipCode     uint   `json:"zip_code,string" form:"zip_code" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	StoreName   string `json:"store_name" form:"store_name"`
	StoreStatus bool   `json:"store_status,string" form:"store_status" gorm:"default:0"`
	OtpCode     string `json:"otp_code" form:"otp_code"`
	IsActive    bool   `json:"is_active,string" form:"is_active"`
	IsAdmin     bool   `json:"is_admin,string" form:"is_admin" `
	Token       string `json:"token" form:"token"`
}

// type UserResponse struct {
// 	Name        string `json:"name" form:"name" validate:"required"`
// 	Email       string `json:"email" form:"email" validate:"required,email"`
// 	Password    string `json:"password" form:"password" validate:"required"`
// 	CityId      uint   `json:"city_id,string" form:"city_id" validate:"required" `
// 	ZipCode     uint   `json:"zip_code,string" form:"zip_code" validate:"required"`
// 	Address     string `json:"address" form:"address" validate:"required"`
// 	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
// 	StoreName   string `json:"store_name" form:"store_name"`
// 	StoreStatus bool   `json:"store_status,string" form:"store_status" gorm:"default:0"`
// 	OtpCode     string `json:"otp_code" form:"otp_code"`
// 	IsActive    bool   `json:"is_active,string" form:"is_active"`
// 	IsAdmin     bool   `json:"is_admin,string" form:"is_admin" `
// 	Token       string `json:"token" form:"token"`
// }

type UserLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
