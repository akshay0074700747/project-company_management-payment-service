package adapters

import (
	"github.com/akshay0074700747/Project/entities"
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

func (payment *PaymentAdapter) AddSubscriptionPlans(req entities.Subscriptions) error {

	query := "INSERT INTO subscriptions (subscription_id,price,description,plan) VALUES($1,$2,$3,$4)"

	if err := payment.DB.Exec(query, req.SubscriptionID, req.Price, req.Description, req.Plan).Error; err != nil {
		return err
	}

	return nil
}

func (payment *PaymentAdapter) GetSubscriptions() ([]entities.Subscriptions, error) {

	query := "SELECT * FROM subscriptions"
	var res []entities.Subscriptions

	if err := payment.DB.Raw(query).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
