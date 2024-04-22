package handlers

import (
	"barber/database"
	"barber/models"
)

func FindById(id uint64) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
