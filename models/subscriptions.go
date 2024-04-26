package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	Base
	Status    string `json:"status"`
	PriceID   string `json:"priceId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint `json:"-"`
}

type AccountSubscription struct {
	ID     uuid.UUID `json:"id,omitempty"`
	Status string    `json:"status,omitempty"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}
