package controllers

import (
	handlers "barber/handlers/users"
	helpers "barber/utils"
	"net/http"
)

func CheckSubscription(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	subscriptionStatus, err := handlers.CheckSubscription(userID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, subscriptionStatus)
}
