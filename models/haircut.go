package models

import (
	"errors"

	"gorm.io/gorm"
)

type Haircut struct {
	gorm.Model
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status bool    `json:"status" gorm:"default:true"`
	UserID uint64  `json:"userId" gorm:"foreignkey:UserID"`
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
