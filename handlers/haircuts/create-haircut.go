package handlers

import (
	"barber/database"
	"barber/models"
)

func CreateHaircut(haircut models.Haircut) (models.Haircut, error) {
	result := database.DB.Create(&haircut)
	if result.Error != nil {
		return models.Haircut{}, result.Error
	}

	return haircut, nil
}
