package entities

import "time"

type PaymentDetails struct {
	PaymentID  string `gorm:"primaryKey"`
	OrderID    string `gorm:"foreignKey:OrderID;references:orders(order_id)"`
	PaymentRef string
	UpdatedAt  time.Time
}

type Subscriptions struct {
	SubscriptionID string `gorm:"primaryKey"`
	Price          uint
	Description    string
	Plan           string
}

type Orders struct {
	OrderID        string `gorm:"primaryKey"`
	UserID         string
	AssetID        string
	SubscriptionID string `gorm:"foreignKey:SubscriptionID;references:subscriptions(subscription_id)"`
	IsPayed        bool
}
