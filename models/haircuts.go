package models

import (
	"errors"

	"github.com/google/uuid"
)

type Haircut struct {
	Base
	Name   string    `json:"name"`
	Price  float64   `json:"price"`
	Status bool      `json:"status" gorm:"default:true"`
	UserID uuid.UUID `json:"userId" gorm:"foreignkey:UserID"`
}

func (Haircut) TableName() string {
	return "haircuts"
}

func (haircut *Haircut) Validate() error {
	if haircut.Name == "" {
		return errors.New("nome é obrigatório")
	}
	if haircut.Price == 0 {
		return errors.New("preço é obrigatório")
	}

	return nil
}
