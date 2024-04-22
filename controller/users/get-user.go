package controllers

import (
	handlers "barber/handlers/users"
	helpers "barber/utils"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	user, err := handlers.FindById(userID)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err)
		return
	}

	accountUser := user.GenerateAccountUser()

	helpers.JSON(w, http.StatusCreated, accountUser)
}
