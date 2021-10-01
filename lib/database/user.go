package database

import (
	"errors"
	"fmt"
	"go-ecommerce/config"
	"go-ecommerce/lib/utils"
	"go-ecommerce/middlewares"
	"go-ecommerce/models"

	"golang.org/x/crypto/bcrypt"
)

var users []models.Users

func GetUsers() (interface{}, error) {
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func StoreUser(user *models.Users) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

func FindUserByEmail(email string) (interface{}, error) {
	var user models.UserResponse
	e := config.DB.Table("users").Select("users.id, users.name, users.email,users.name, users.provinces_id, provinces.name as provinces_name, users.city_id, cities.name as city_name, users.zip_code, users.address, users.phone_number, users.store_status, users.otp_code, users.is_active, users.is_admin, users.token ").Joins("join provinces on provinces.id = users.provinces_id").Joins("join cities on cities.id = users.city_id").Where("users.email = ?", email).First(&user).Error
	if e != nil {
		return nil, e
	}
	return user, nil
}

func FindUserByPhone(phone string) (interface{}, error) {
	var user models.UserResponse
	e := config.DB.Table("users").Select("users.id, users.name, users.email,users.name, users.provinces_id, provinces.name as provinces_name, users.city_id, cities.name as city_name, users.zip_code, users.address, users.phone_number, users.store_status, users.otp_code, users.is_active, users.is_admin, users.token ").Joins("join provinces on provinces.id = users.provinces_id").Joins("join cities on cities.id = users.city_id").Where("users.phone_number = ?", phone).First(&user).Error

	if e != nil {
		return nil, e
	}
	return user, nil
}

func FindUserById(userId int) (interface{}, error) {
	var user models.UserResponse
	e := config.DB.Table("users").Select("users.id, users.name, users.email,users.name, users.provinces_id, provinces.name as provinces_name, users.city_id, cities.name as city_name, users.zip_code, users.address, users.phone_number, users.store_status, users.otp_code, users.is_active, users.is_admin, users.token ").Joins("join provinces on provinces.id = users.provinces_id").Joins("join cities on cities.id = users.city_id").Where("users.id = ?", userId).First(&user).Error

	if e != nil {
		return nil, e
	}

	return user, nil
}

func LoginUser(email string, hasPass string) (interface{}, error) {
	var err error
	var user models.Users

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(hasPass))
	if errPass != nil {
		return nil, errPass
	}
	id := int(user.ID)
	user.Token, err = middlewares.CreateToken(id, user.IsAdmin, user.IsActive)
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	// jika findEmail kosong maka

	return user, nil

}

// re create otp
func ReCreateOTP(id int, otp, text string) (interface{}, error) {
	var user models.Users

	e := config.DB.Where("id=?", id).First(&user).Error
	if e != nil {
		return nil, e
	}
	// fmt.Println(e)
	user.OtpCode = otp
	fmt.Println(user.PhoneNumber)
	// send to otp
	wa, e := utils.SendWa(string(user.PhoneNumber), text)
	if e != nil {
		fmt.Println("Error")
		fmt.Println(e)
	}
	fmt.Println("WA")
	fmt.Println(wa)

	config.DB.Save(&user)
	return user, nil
}

func VerifOTP(id int, requestotp string) (interface{}, error) {
	var user models.Users
	var err error
	e := config.DB.Where("id=?", id).First(&user).Error
	if e != nil {
		return nil, e
	}

	if user.OtpCode != requestotp {
		return nil, errors.New("OTP code doesn't same")
	}
	// jika sama maka otp dihapus
	user.OtpCode = ""
	user.IsActive = true
	user.Token, err = middlewares.CreateToken(id, user.IsAdmin, user.IsActive)
	if err != nil {
		return nil, err
	}

	config.DB.Save(&user)

	return user, nil

}
