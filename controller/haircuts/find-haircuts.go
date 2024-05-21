package controllers

import (
	handlers "barber/handlers/haircuts"
	helpers "barber/utils"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func FindHaircut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	haircutIDStr := vars["ID"]

	haircutID, err := uuid.Parse(haircutIDStr)
	if err != nil {
		fmt.Println("Erro ao converter string para UUID:", err)
		return
	}

	haircut, err := handlers.FindById(haircutID)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, haircut)
}
