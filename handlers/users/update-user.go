package handlers

import (
	"barber/database"
	"barber/models"
)

func UpdateUser(user models.User) (models.User, error) {
	result := database.DB.Save(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
