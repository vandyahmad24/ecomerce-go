package database

import (
	"go-ecommerce/config"
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

func FindUserByEmail(user *models.Users, email string) (interface{}, error) {
	e := config.DB.Where("email=?", email).First(&users).Error
	if e != nil {
		return nil, e
	}
	return user, nil
}

func FindUserByPhone(user *models.Users, phone string) (interface{}, error) {
	e := config.DB.Where("phone_number=?", phone).First(&users).Error
	if e != nil {
		return nil, e
	}
	return user, nil
}

func FindUserById(userId int) (interface{}, error) {
	var user models.Users
	if e := config.DB.Find(&user, userId).Error; e != nil {
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

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	// jika findEmail kosong maka

	return user, nil

}
