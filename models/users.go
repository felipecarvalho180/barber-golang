package models

import (
	helpers "barber/utils"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string       `json:"name"`
	Email          string       `json:"email" gorm:"uniqueIndex"`
	Address        string       `json:"address,omitempty"`
	Password       string       `json:"password"`
	Token          string       `json:"token,omitempty" gorm:"-"`
	SubscriptionID uint         `json:"-"`
	Subscription   Subscription `json:"subscription,omitempty" gorm:"foreignkey:UserID"`
	Haircuts       []Haircut    `json:"haircuts,omitempty" gorm:"foreignkey:UserID"`
	Service        []Service    `json:"service,omitempty" gorm:"foreignkey:HaircutID"`
}

type AccountUser struct {
	ID           uint                 `json:"id"`
	Name         string               `json:"name"`
	Email        string               `json:"email"`
	Address      string               `json:"address,omitempty"`
	Token        string               `json:"token,omitempty"`
	Subscription *AccountSubscription `json:"subscription,omitempty"`
}

func (User) TableName() string {
	return "users"
}

const (
	SIGN_UP_STEP = iota
	UPDATE_USER_STEP
	UPDATE_PET_STEP
)

func (user *User) validate(step int) error {
	if user.Name == "" {
		return errors.New("nome é obrigatório")
	}
	if user.Email == "" {
		return errors.New("email é obrigatório")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("formato de email inválido")
	}
	if step == SIGN_UP_STEP && user.Password == "" {
		return errors.New("senha é obrigatória")
	}

	return nil
}

func (user *User) format(step int) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if step == SIGN_UP_STEP {
		passwordWithHash, err := helpers.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordWithHash)
	}

	return nil
}

func (user *User) Prepare(step int) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) GenerateAccountUser() AccountUser {
	var subscription *AccountSubscription
	if user.Subscription.ID != 0 {
		subscription = &AccountSubscription{
			ID:     user.Subscription.ID,
			Status: user.Subscription.Status,
		}
	} else {
		subscription = nil
	}

	var accountUser = AccountUser{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Address:      user.Address,
		Subscription: subscription,
		Token:        user.Token,
	}

	return accountUser
}