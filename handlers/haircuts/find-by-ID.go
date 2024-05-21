package handlers

import (
	"barber/database"
	"barber/models"

	"github.com/google/uuid"
)

func FindById(id uuid.UUID) (models.Haircut, error) {
	var haircut models.Haircut
	result := database.DB.First(&haircut, id)
	if result.Error != nil {
		return models.Haircut{}, result.Error
	}

	return haircut, nil
}
