package controllers

import (
	handlers "barber/handlers/haircuts"
	userHandlers "barber/handlers/users"
	"barber/models"
	helpers "barber/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func UpdateHaircut(w http.ResponseWriter, r *http.Request) {
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

	if err = haircut.UpdateValidate(); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	userDB, err := userHandlers.FindById(userID)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err)
		return
	}

	if userDB.SubscriptionID == uuid.Nil {
		helpers.Error(w, http.StatusForbidden, errors.New("assinatura necessária para executar essa operação"))
		return
	}

	vars := mux.Vars(r)
	haircutIDStr := vars["ID"]

	haircutID, err := uuid.Parse(haircutIDStr)
	if err != nil {
		fmt.Println("Erro ao converter string para UUID:", err)
		return
	}

	haircutDB, err := handlers.FindById(haircutID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	if haircutDB.UserID != userID {
		helpers.Error(w, http.StatusUnauthorized, errors.New("usuário inválido"))
		return
	}

	haircutDB.Name = haircut.Name
	haircutDB.Price = haircut.Price
	haircutDB.Status = haircut.Status

	updatedHaircut, err := handlers.UpdateHaircut(haircutDB)
	if err != nil {
		helpers.Error(w, http.StatusNotFound, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, updatedHaircut)
}
