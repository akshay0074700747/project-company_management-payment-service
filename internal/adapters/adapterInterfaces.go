package adapters

import "github.com/akshay0074700747/Project/entities"

type PaymentAdapterInterfaces interface {
	AddSubscriptionPlans(entities.Subscriptions)(error)
	GetSubscriptions()([]entities.Subscriptions,error)
	Subcribe(entities.Orders)(entities.Orders,error)
	GetallSubscriptions()([]entities.Orders,error)
	GetOrderDetails(string)(entities.MakePaymentUsecase,error)
	MakePayment(entities.PaymentDetails)(error)
}
