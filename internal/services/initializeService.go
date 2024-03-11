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

	// r.GET("/project/snapshots", engine.Srv.GetSnapshots)
	r.GET("/project/task/stages", engine.Srv.getStages)

	r.LoadHTMLGlob("payment.html")

	http.ListenAndServe(addr, handler)
}