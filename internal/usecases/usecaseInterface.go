package usecases

import "github.com/akshay0074700747/Project/entities"

type PaymentUsecaseInterfaces interface {
	AddSubscriptionPlans(req entities.Subscriptions) error
	GetSubscriptions()([]entities.Subscriptions,error)
}
