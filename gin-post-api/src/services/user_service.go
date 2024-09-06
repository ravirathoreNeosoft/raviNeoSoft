package services

import (
	"errors"
	"gin-post-api/src/config"
	"gin-post-api/src/models"
)

func CreateUser(user models.User) (models.User, error) {

	if user.Age < 18 {
		return models.User{}, errors.New("user must be at least 18 years old")
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
