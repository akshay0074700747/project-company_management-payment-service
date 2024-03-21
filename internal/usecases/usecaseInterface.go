package usecases

import "github.com/akshay0074700747/Project/entities"

type PaymentUsecaseInterfaces interface {
	AddSubscriptionPlans(req entities.Subscriptions) error
	GetSubscriptions()([]entities.Subscriptions,error)
	Subcribe(entities.Orders)(entities.Orders,error)
	GetallSubscriptions()([]entities.Orders,error)
	GetOrderDetails(string)(entities.MakePaymentUsecase,error)
	MakePayment(entities.PaymentDetails)(error)
	GetPaymentDetailsofUser(string)([]entities.PaymentDetails,error)
	VerifyTransaction(string,string)(bool,error)
	UpdateAsset(entities.UpdateAssetID)(error)
	GetAssetID(string)(bool,error)
	GetAssetIDfromOrderID(string)(string,error)
}
