package controllers

import (
	handlers "barber/handlers/haircuts"
	helpers "barber/utils"
	"errors"
	"net/http"
	"strconv"
)

func GetHaircuts(w http.ResponseWriter, r *http.Request) {
	statusStr := r.URL.Query().Get("status")
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, errors.New("status inválido"))
		return
	}

	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	haircuts, err := handlers.GetHaircuts(userID, status)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, haircuts)
}
