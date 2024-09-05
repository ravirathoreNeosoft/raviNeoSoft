package services

import (
	"errors"
	"gin-post-api/src/models"
)

func CreateUser(user models.User) (models.User, error) {

	if user.Age < 18 {
		return models.User{}, errors.New("can't create user! as the age must be at least 18 years old")
	}

	return user, nil
}
