package controllers

import (
	haircutHandlers "barber/handlers/haircuts"
	handlers "barber/handlers/users"
	"barber/models"
	helpers "barber/utils"
	"encoding/json"
	"io"
	"net/http"
)

func CreateHaircut(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var haircut models.Haircut
	if err = json.Unmarshal(body, &haircut); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = haircut.Validate(); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if _, err = handlers.FindById(userID); err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	haircut.UserID = userID

	haircutDB, err := haircutHandlers.CreateHaircut(haircut)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, haircutDB)
}
