package models

import "time"

type Subscription struct {
	ID        uint   `gorm:"primarykey"`
	Status    string `json:"status"`
	PriceID   string `json:"priceId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint `json:"-"`
}

type AccountSubscription struct {
	ID     uint   `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}
