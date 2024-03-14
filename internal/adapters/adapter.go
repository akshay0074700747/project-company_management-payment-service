package adapters

import (
	"time"

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

func (payment *PaymentAdapter) Subcribe(req entities.Orders) (entities.Orders, error) {

	query := "INSERT INTO orders (order_id,user_id,asset_id,subscription_id) VALUES($1,$2,$3,$4) RETURNING order_id,user_id,assett_id,subscription_id,is_payed"
	var res entities.Orders
	if err := payment.DB.Raw(query, req.OrderID, req.UserID, req.AssetID, req.SubscriptionID).Scan(&res).Error; err != nil {
		return entities.Orders{}, err
	}

	return res, nil
}

func (payment *PaymentAdapter) GetallSubscriptions() ([]entities.Orders, error) {

	query := "SELECT * FROM orders"
	var res []entities.Orders

	if err := payment.DB.Raw(query).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (payment *PaymentAdapter) GetOrderDetails(orderID string) (entities.MakePaymentUsecase, error) {

	var res entities.MakePaymentUsecase
	query := "SELECT o.order_id,o.user_id,o.is_payed,s.price FROM orders o INNER JOIN subscriptions s ON o.subscription_id = s.subscription_id WHERE order_id = $1 "

	if err := payment.DB.Raw(query).Scan(&res).Error; err != nil {
		return entities.MakePaymentUsecase{}, err
	}

	return res, nil
}

func (snap *PaymentAdapter) MakePayment(req entities.PaymentDetails) error {

	tx := snap.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	query := "INSERT INTO payment_details (payment_id,order_id,payment_ref,updated_at) VALUES($1,$2,$3,$4)"
	if err := tx.Exec(query, req.PaymentID, req.OrderID, req.PaymentRef, time.Now()).Error; err != nil {
		tx.Rollback()
		return err
	}

	query = "UPDATE orders SET is_payed = true WHERE order_id = $1"
	if err := tx.Exec(query, req.OrderID).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (snap *PaymentAdapter) GetPaymentDetailsofUser(userID string) ([]entities.PaymentDetails, error) {

	query := "SELECT * FROM payment_details WHERE order_id IN (select order_id from orders WHERE user_id = $1)"
	var res []entities.PaymentDetails

	if err := snap.DB.Raw(query, userID).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (snap *PaymentAdapter) VerifyTransaction(transactionID, userID string) (bool, error) {

	query := "SELECT o.asset_id FROM payment_details p INNER JOIN orders o ON p.order_id = o.order_id AND o.user_id = $1 WHERE payment_id = $2"
	var res string

	if err := snap.DB.Raw(query, userID, transactionID).Scan(&res).Error; err != nil {
		return false, err
	}

	if res != "" {
		return false, nil
	}

	return true, nil
}

func (snap *PaymentAdapter) UpdateAsset(req entities.UpdateAssetID) error {

	query := "UPDATE orders SET asset_id = $1 WHERE order_id = (SELECT order_id FROM payment_details WHERE payment_id = $2) AND user_id = $3"

	if err := snap.DB.Exec(query, req.AssetID, req.TransactionID, req.UserID).Error; err != nil {
		return err
	}

	return nil
}

func (snap *PaymentAdapter) GetAssetID(assetID string) (bool, error) {

	query := "SELECT * FROM orders WHERE asset_id = $1 AND is_payed = true"

	tx := snap.DB.Raw(query, assetID)
	if tx.Error != nil {
		return false, tx.Error
	}

	if tx.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
