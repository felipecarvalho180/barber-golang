package handlers

import (
	"barber/database"
	"barber/models"
)

func UpdateHaircut(haircut models.Haircut) (models.Haircut, error) {
	result := database.DB.Save(&haircut)
	if result.Error != nil {
		return models.Haircut{}, result.Error
	}

	return haircut, nil
}
