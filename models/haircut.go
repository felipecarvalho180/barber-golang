package models

import (
	"gorm.io/gorm"
)

type Haircut struct {
	gorm.Model
	Name    string    `json:"name"`
	Price   string    `json:"price"`
	Status  bool      `json:"status" gorm:"default:true"`
	UserID  uint      `json:"-"`
	User    User      `json:"user,omitempty"`
	Service []Service `json:"service,omitempty" gorm:"foreignkey:HaircutID"`
}

func (Haircut) TableName() string {
	return "haircuts"
}
