package models

import (
	"time"
)

type Subscription struct {
	Base
	Status    string `json:"status"`
	PriceID   string `json:"priceId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Subscription) TableName() string {
	return "subscriptions"
}
