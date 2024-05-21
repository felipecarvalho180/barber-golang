package models

import (
	helpers "barber/utils"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
)

type User struct {
	Base
	Name           string    `json:"name"`
	Email          string    `json:"email" gorm:"uniqueIndex"`
	Address        string    `json:"address,omitempty"`
	Password       string    `json:"password"`
	Token          string    `json:"token,omitempty" gorm:"-"`
	SubscriptionID uuid.UUID `json:"-"`
	Haircuts       []Haircut `json:"haircuts,omitempty" gorm:"foreignkey:UserID"`
	Service        []Service `json:"service,omitempty" gorm:"foreignkey:HaircutID"`
}

type AccountUser struct {
	ID           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Address      string     `json:"address,omitempty"`
	Token        string     `json:"token,omitempty"`
	Subscription *uuid.UUID `json:"subscription,omitempty"`
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
	var subscription *uuid.UUID
	if user.SubscriptionID != uuid.Nil {
		subscription = &user.SubscriptionID
	} else {
		subscription = nil
	}

	var accountUser = AccountUser{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Address:      user.Address,
		Token:        user.Token,
		Subscription: subscription,
	}

	return accountUser
}
