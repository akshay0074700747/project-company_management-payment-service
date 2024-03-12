package entities

import "time"

type PaymentDetails struct {
	PaymentID  string    `gorm:"primaryKey" json:"PaymentID"`
	OrderID    string    `gorm:"foreignKey:OrderID;references:orders(order_id)" json:"OrderID"`
	PaymentRef string    `json:"PaymentRef"`
	UpdatedAt  time.Time `json:"UpdatedAt" `
}

type Subscriptions struct {
	SubscriptionID string `gorm:"primaryKey" json:"subscription_id"`
	Price          uint   `json:"price"`
	Description    string `json:"description"`
	Plan           string `json:"plan"`
}

type Orders struct {
	OrderID        string `gorm:"primaryKey" json:"OrderID"`
	UserID         string `json:"UserID"`
	AssetID        string `json:"AssetID"`
	SubscriptionID string `gorm:"foreignKey:SubscriptionID;references:subscriptions(subscription_id)" json:"SubscriptionID"`
	IsPayed        bool   `gorm:"default:false" json:"IsPayed"`
}
