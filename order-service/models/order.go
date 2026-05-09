package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	StatusPending    OrderStatus = "PENDING"
	StatusProcessing OrderStatus = "PROCESSING"
	StatusExecuted   OrderStatus = "EXECUTED"
	StatusFailed     OrderStatus = "FAILED"
)

type OrderType string

const (
	TypeSIP     OrderType = "SIP"
	TypeLumpsum OrderType = "LUMPSUM"
)

type Order struct {
	ID            string      `gorm:"primaryKey" json:"id"`
	UserID        string      `gorm:"index" json:"user_id"`
	FundID        string      `json:"fund_id"`
	Amount        float64     `json:"amount"`
	Type          OrderType   `json:"type"`
	Status        OrderStatus `json:"status"`
	IdempotencyKey string      `gorm:"uniqueIndex" json:"idempotency_key"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}
	return
}
