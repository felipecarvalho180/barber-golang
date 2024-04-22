package controllers

import (
	handlers "barber/handlers/users"
	"barber/models"
	helpers "barber/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	userDB, err := handlers.FindByEmail(user.Email)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, errors.New("credenciais inválidas"))
		return
	}

	if err = helpers.VerifyPassword(user.Password, userDB.Password); err != nil {
		helpers.Error(w, http.StatusUnauthorized, errors.New("credenciais inválidas"))
		return
	}

	token, err := helpers.GenerateToken(uint64(userDB.ID))
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	userDB.Token = token
	accountUser := userDB.GenerateAccountUser()

	helpers.JSON(w, http.StatusOK, accountUser)
}
