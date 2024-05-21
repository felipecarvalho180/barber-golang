package handlers

import (
	"barber/database"
	"barber/models"

	"github.com/google/uuid"
)

func CheckSubscription(id uuid.UUID) (bool, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return false, result.Error
	}

	subscriptionStatus := false
	if user.SubscriptionID != uuid.Nil {
		subscriptionStatus = true
	}

	return subscriptionStatus, nil
}
