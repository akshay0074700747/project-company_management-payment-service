package injectdependency

import (
	"github.com/akshay0074700747/Project/config"
	"github.com/akshay0074700747/Project/db"
	"github.com/akshay0074700747/Project/internal/adapters"
	"github.com/akshay0074700747/Project/internal/services"
	"github.com/akshay0074700747/Project/internal/usecases"
)

func Initialize(cfg config.Config) *services.PaymentEngine {

	db := db.ConnectDB(cfg)
	adapter := adapters.NewPaymentAdapter(db)
	usecase := usecases.NewPaymentUseCases(adapter)
	service := services.NewPaymentService(usecase)

	return services.NewPaymentEngine(service)
}
