package handlers

import (
	"barber/database"
	"barber/models"

	"github.com/google/uuid"
)

func FindById(id uuid.UUID) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
