package adapters

import "github.com/akshay0074700747/Project/entities"

type PaymentAdapterInterfaces interface {
	AddSubscriptionPlans(entities.Subscriptions)(error)
	GetSubscriptions()([]entities.Subscriptions,error)
}
