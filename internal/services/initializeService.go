package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type PaymentEngine struct {
	Srv *PaymentService
}

func NewPaymentEngine(srv *PaymentService) *PaymentEngine {
	return &PaymentEngine{
		Srv: srv,
	}
}
func (engine *PaymentEngine) Start(addr string) {

	r := gin.New()
	r.Use(gin.Logger())

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	r.POST("/subscription/plan/add", engine.Srv.AddSubscriptionPlan)
	r.GET("/subscription/plan", engine.Srv.GetallSubscriptionPlans)
	r.POST("/subscription/plan/subscribe", engine.Srv.subscribe)
	r.GET("/subscriptions", engine.Srv.getSubscriptions)
	r.GET("/subscription/plan/subscribe/order/pay", engine.Srv.subscriptionPayment)
	r.GET("/verify/payment", engine.Srv.verifyPayment)
	r.GET("/payment/verified", engine.Srv.servePaymentSuccesspage)
	r.GET("/payments", engine.Srv.payments)
	r.GET("/transaction/verify", engine.Srv.verifyTransaction)
	r.PATCH("/transaction/update", engine.Srv.updateAssetID)
	r.GET("/transaction/project", engine.Srv.getAssetID)

	r.LoadHTMLGlob("*.html")

	http.ListenAndServe(addr, handler)
}
