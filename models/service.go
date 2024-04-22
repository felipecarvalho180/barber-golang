package models

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Customer string `json:"name"`

	HaircutID uint    `json:"-"`
	Haircut   Haircut `json:"haircut,omitempty"`

	UserID uint `json:"-"`
	User   User `json:"user,omitempty"`
}

func (Service) TableName() string {
	return "services"
}
