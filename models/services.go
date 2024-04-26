package models

import (
	"github.com/google/uuid"
)

type Service struct {
	Base
	Customer string `json:"name"`

	HaircutID uuid.UUID `json:"-"`
	Haircut   Haircut   `json:"haircut,omitempty"`

	UserID uuid.UUID `json:"-"`
	User   User      `json:"user,omitempty"`
}

func (Service) TableName() string {
	return "services"
}
