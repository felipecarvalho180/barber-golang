package controllers

import (
	handlers "barber/handlers/users"
	"barber/models"
	helpers "barber/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare(models.SIGN_UP_STEP); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	user, err = handlers.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			helpers.Error(w, http.StatusConflict, errors.New("o e-mail já está em uso"))
			return
		}

		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	token, err := helpers.GenerateToken(uint64(user.ID))
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.Token = token
	accountUser := user.GenerateAccountUser()

	helpers.JSON(w, http.StatusCreated, accountUser)
}
