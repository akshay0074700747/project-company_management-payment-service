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
