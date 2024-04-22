package handlers

import (
	"barber/database"
	"barber/models"
)

func CreateUser(user models.User) (models.User, error) {
	result := database.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
