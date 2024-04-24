package controllers

import (
	handlers "barber/handlers/users"
	"barber/models"
	helpers "barber/utils"
	"encoding/json"
	"io"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	userDB, err := handlers.FindById(userID)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	userDB.Name = user.Name
	userDB.Address = user.Address

	updatedUser, err := handlers.UpdateUser(userDB)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err)
		return
	}

	accountUser := updatedUser.GenerateAccountUser()

	helpers.JSON(w, http.StatusCreated, accountUser)
}
