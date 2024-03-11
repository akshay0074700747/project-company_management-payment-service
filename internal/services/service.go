package services

import (
	"github.com/akshay0074700747/Project/internal/usecases"
	"github.com/gin-gonic/gin"
)

type PaymentService struct {
	Usecase usecases.PaymentUsecaseInterfaces
}

func NewPaymentService(usecase usecases.PaymentUsecaseInterfaces) *PaymentService {
	return &PaymentService{
		Usecase: usecase,
	}
}

func (snap *PaymentService) getStages(c *gin.Context) {

	c.HTML(200, "payment.html", gin.H{})
}
