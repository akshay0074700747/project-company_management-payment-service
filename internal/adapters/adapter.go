package adapters

import (
	"gorm.io/gorm"
)

type PaymentAdapter struct {
	DB *gorm.DB
}

func NewPaymentAdapter(db *gorm.DB) *PaymentAdapter {
	return &PaymentAdapter{
		DB: db,
	}
}
