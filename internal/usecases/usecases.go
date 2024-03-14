package usecases

import (
	"github.com/akshay0074700747/Project/entities"
	"github.com/akshay0074700747/Project/helpers"
	"github.com/akshay0074700747/Project/internal/adapters"
)

type PaymentUseCases struct {
	Adapter adapters.PaymentAdapterInterfaces
}

func NewPaymentUseCases(adapter adapters.PaymentAdapterInterfaces) *PaymentUseCases {
	return &PaymentUseCases{
		Adapter: adapter,
	}
}

func (payment *PaymentUseCases) AddSubscriptionPlans(req entities.Subscriptions) error {

	req.SubscriptionID = helpers.GenUuid()

	if err := payment.Adapter.AddSubscriptionPlans(req); err != nil {
		helpers.PrintErr(err, "error happened at AddSubscriptionPlans adapter")
		return err
	}

	return nil
}

func (payment *PaymentUseCases) GetSubscriptions() ([]entities.Subscriptions, error) {

	res, err := payment.Adapter.GetSubscriptions()
	if err != nil {
		helpers.PrintErr(err, "erorr happened at GetSubscriptions adapter")
		return nil, err
	}

	return res, nil
}

func (payment *PaymentUseCases) Subcribe(req entities.Orders) (entities.Orders, error) {

	req.OrderID = helpers.GenUuid()
	res, err := payment.Adapter.Subcribe(req)
	if err != nil {
		helpers.PrintErr(err, "erroor happened at Subcribe adapter")
		return entities.Orders{}, err
	}

	return res, nil

}

func (payment *PaymentUseCases) GetallSubscriptions() ([]entities.Orders, error) {

	res, err := payment.Adapter.GetallSubscriptions()
	if err != nil {
		helpers.PrintErr(err, "error happened at GetallSubscriptions adapter")
		return nil, err
	}

	return res, nil
}

func (payment *PaymentUseCases) GetOrderDetails(orderID string) (entities.MakePaymentUsecase, error) {

	res, err := payment.Adapter.GetOrderDetails(orderID)
	if err != nil {
		helpers.PrintErr(err, "errror happened at GetOrderDetails adapter")
		return entities.MakePaymentUsecase{}, err
	}

	return res, nil
}

func (payment *PaymentUseCases) MakePayment(req entities.PaymentDetails) error {

	req.PaymentID = helpers.GenUuid()
	if err := payment.Adapter.MakePayment(req); err != nil {
		helpers.PrintErr(err, "errorr happened at MakePayment adapter")
		return err
	}

	return nil
}

func (snap *PaymentUseCases) GetPaymentDetailsofUser(userID string) ([]entities.PaymentDetails, error) {

	res, err := snap.Adapter.GetPaymentDetailsofUser(userID)
	if err != nil {
		helpers.PrintErr(err, "error happened at GetPaymentDetailsofUser adapter")
		return nil, err
	}

	return res, nil
}

func (snap *PaymentUseCases) VerifyTransaction(transactionID, userID string) (bool, error) {

	res, err := snap.Adapter.VerifyTransaction(transactionID, userID)
	if err != nil {
		helpers.PrintErr(err, "error happened at VerifyTransaction adapter")
		return false, err
	}

	return res, nil
}

func (snap *PaymentUseCases) UpdateAsset(req entities.UpdateAssetID) error {

	if err := snap.Adapter.UpdateAsset(req); err != nil {
		helpers.PrintErr(err, "error happened at UpdateAsset usecase")
		return err
	}

	return nil
}

func (snap *PaymentUseCases) GetAssetID(assetID string) (bool, error) {

	res, err := snap.Adapter.GetAssetID(assetID)
	if err != nil {
		helpers.PrintErr(err, "errro happeedn at GetAssetID adapter")
		return false, err
	}

	return res, nil
}
