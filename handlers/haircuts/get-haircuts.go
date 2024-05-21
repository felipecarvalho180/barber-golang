package handlers

import (
	"barber/database"
	"barber/models"

	"github.com/google/uuid"
)

func GetHaircuts(userID uuid.UUID, status bool) ([]models.Haircut, error) {
	var haircuts = []models.Haircut{}
	result := database.DB.Where("user_id = ?", userID).Where("status = ?", status).Find(&haircuts)
	if result.Error != nil {
		return []models.Haircut{}, result.Error
	}

	return haircuts, nil
}
