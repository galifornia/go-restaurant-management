package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	UUID           string        `json:"uuid"`
	OrderID        string        `json:"order_id"`
	PaymentMethod  PaymentMethod `json:"payment_method"`
	PaymentStatus  PaymentStatus `json:"payment_status"`
	PaymentDueDate time.Time     `json:"payment_due_time"`
}

type PaymentMethod int64

const (
	CARD PaymentMethod = 0
	CASH               = 1
)

type PaymentStatus int64

const (
	PENDING PaymentStatus = 0
	PAID                  = 1
)
