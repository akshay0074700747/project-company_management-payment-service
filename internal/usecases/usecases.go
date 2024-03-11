package usecases

import "github.com/akshay0074700747/Project/internal/adapters"

type PaymentUseCases struct {
	Adapter adapters.PaymentAdapterInterfaces
}

func NewPaymentUseCases(adapter adapters.PaymentAdapterInterfaces) *PaymentUseCases {
	return &PaymentUseCases{
		Adapter: adapter,
	}
}
