package services

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/akshay0074700747/Project/config"
	"github.com/akshay0074700747/Project/entities"
	"github.com/akshay0074700747/Project/helpers"
	"github.com/akshay0074700747/Project/internal/usecases"
	"github.com/akshay0074700747/projectandCompany_management_protofiles/pb/companypb"
	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)

type PaymentService struct {
	Usecase  usecases.PaymentUsecaseInterfaces
	Cfg      config.Config
	CompConn companypb.CompanyServiceClient
}

func NewPaymentService(usecase usecases.PaymentUsecaseInterfaces, cfg config.Config) *PaymentService {

	client, _ := helpers.DialGrpc(":50003")

	return &PaymentService{
		Usecase:  usecase,
		Cfg:      cfg,
		CompConn: companypb.NewCompanyServiceClient(client),
	}
}

func (snap *PaymentService) AddSubscriptionPlan(c *gin.Context) {

	var sub entities.Subscriptions
	if err := c.BindJSON(&sub); err != nil {
		helpers.PrintErr(err, "error happened at binding subscriptions")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusOK, entities.Responce{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		})
	}

	if err := snap.Usecase.AddSubscriptionPlans(sub); err != nil {
		helpers.PrintErr(err, "error happened at AddSubscriptionPlans usecase")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusOK, entities.Responce{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		})
	}

	c.JSON(http.StatusOK, entities.Responce{
		StatusCode: 200,
		Message:    "successfully added new subscription plans",
		Error:      nil,
	})
}

func (snap *PaymentService) GetallSubscriptionPlans(c *gin.Context) {

	res, err := snap.Usecase.GetSubscriptions()
	if err != nil {
		helpers.PrintErr(err, "error happened at GetSubscriptions usecase")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusOK, entities.Responce{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		})
	}

	c.JSON(http.StatusOK, entities.Responce{
		StatusCode: 200,
		Message:    "successfully got all subscription plans",
		Data:       res,
	})

}

func (snap *PaymentService) subscribe(c *gin.Context) {

	var sub entities.Orders

	if err := c.BindJSON(&sub); err != nil {
		helpers.PrintErr(err, "error happened at binding subscriptions")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusOK, entities.Responce{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		})
	}

	sub.UserID = c.Request.Header.Get("X-User-ID")

	fmt.Println(sub.UserID)

	res, err := snap.Usecase.Subcribe(sub)
	if err != nil {
		helpers.PrintErr(err, "error happened at Subcribe usecase")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusOK, entities.Responce{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		})
	}

	c.JSON(http.StatusOK, res)
}

func (snap *PaymentService) getSubscriptions(c *gin.Context) {

	res, err := snap.Usecase.GetallSubscriptions()
	if err != nil {
		helpers.PrintErr(err, "error happened at GetallSubscriptions usecase")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusOK, entities.Responce{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		})
	}

	c.JSON(http.StatusOK, entities.Responce{
		StatusCode: 200,
		Message:    "successfully got all subscription plans",
		Data:       res,
	})
}

func (snap *PaymentService) subscriptionPayment(c *gin.Context) {

	orderId := c.Query("orderID")

	orderdata, err := snap.Usecase.GetOrderDetails(orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.Responce{
			StatusCode: 500,
			Message:    "cant find data",
			Error:      err,
		})
		return
	}

	if orderdata.IsPayed {
		c.JSON(http.StatusBadRequest, entities.Responce{
			StatusCode: 400,
			Message:    "already payed",
		})
		return
	}

	client := razorpay.NewClient(snap.Cfg.RAZORPAYID, snap.Cfg.RAZORPAYSECRET)

	data := map[string]interface{}{
		"amount":   orderdata.Price * 100,
		"currency": "INR",
		"receipt":  "test_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.Responce{
			StatusCode: 500,
			Message:    "cant process order right now",
			Error:      err,
		})
		return
	}

	value := body["id"]
	razorpayID := value.(string)

	c.HTML(200, "payment.html", gin.H{
		"total_price": orderdata.Price,
		"total":       orderdata.Price,
		"orderData":   orderId,
		"orderid":     razorpayID,
		"amount":      orderdata.Price,
		"userID":      orderdata.UserID,
	})

}

func (snap *PaymentService) verifyPayment(c *gin.Context) {
	paymentRef := c.Query("payment_ref")

	idStr := c.Query("order_id")

	orderID := strings.ReplaceAll(idStr, " ", "")

	if err := snap.Usecase.MakePayment(entities.PaymentDetails{
		OrderID:    orderID,
		PaymentRef: paymentRef,
	}); err != nil {
		helpers.PrintErr(err, "error happened at MakePayment usecase")
		c.JSON(http.StatusInternalServerError, entities.Responce{
			StatusCode: 500,
			Message:    "cant process order right now",
			Error:      err,
		})
		return
	}

	go func() {
		assetID, err := snap.Usecase.GetAssetIDfromOrderID(orderID)
		if err != nil {
			helpers.PrintErr(err, "error happned")
		}
		_, err = snap.CompConn.ToggleIsPayed(context.TODO(), &companypb.ToggleIsPayedReq{
			CompanyID: assetID,
			IsPayed:   true,
		})
		if err != nil {
			helpers.PrintErr(err, "error happened")
		}
	}()

	c.JSON(http.StatusOK, entities.Responce{
		StatusCode: 200,
		Message:    "payment updated",
		Data:       true,
	})
}

func (snap *PaymentService) servePaymentSuccesspage(c *gin.Context) {

	c.HTML(200, "paymentVerified.html", gin.H{})
}

func (snap *PaymentService) payments(c *gin.Context) {

	userID := c.Request.Header.Get("X-User-ID")

	res, err := snap.Usecase.GetPaymentDetailsofUser(userID)
	if err != nil {
		helpers.PrintErr(err, "error happened at GetPaymentDetailsofUser usecase")
		c.JSON(http.StatusInternalServerError, entities.Responce{
			StatusCode: 500,
			Message:    "cant get the payments right now",
			Error:      err,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (snap *PaymentService) verifyTransaction(c *gin.Context) {

	transactionID := c.Query("transactionID")
	userID := c.Query("userID")

	res, err := snap.Usecase.VerifyTransaction(transactionID, userID)
	if err != nil {
		helpers.PrintErr(err, "error happened at VerifyTransaction usecase")
		c.JSON(http.StatusInternalServerError, entities.Responce{
			StatusCode: 500,
			Data:       false,
			Error:      err,
		})
		return
	}

	if !res {
		c.JSON(http.StatusBadRequest, entities.Responce{
			StatusCode: 500,
			Data:       false,
			Error:      nil,
		})
	}

	c.JSON(http.StatusOK, entities.Responce{
		StatusCode: 200,
		Data:       true,
		Error:      nil,
	})
}

func (snap *PaymentService) updateAssetID(c *gin.Context) {

	var sub entities.UpdateAssetID

	if err := c.BindJSON(&sub); err != nil {
		helpers.PrintErr(err, "error happened at binding subscriptions")
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusOK, entities.Responce{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		})
		return
	}

	if err := snap.Usecase.UpdateAsset(sub); err != nil {
		helpers.PrintErr(err, "error happened at UpdateAsset usecase")
		c.JSON(http.StatusInternalServerError, entities.Responce{
			StatusCode: 500,
			Error:      err,
		})
		return
	}

	c.JSON(http.StatusOK, entities.Responce{
		StatusCode: 200,
		Data:       false,
		Error:      nil,
	})
}

func (snap *PaymentService) getAssetID(c *gin.Context) {

	assetID := c.Query("assetID")

	res, err := snap.Usecase.GetAssetID(assetID)
	if err != nil {
		helpers.PrintErr(err, "error happened at GetAssetID usecase")
		c.JSON(http.StatusInternalServerError, entities.Responce{
			StatusCode: 500,
			Error:      err,
		})
		return
	}

	if !res {
		c.JSON(http.StatusInternalServerError, entities.Responce{
			StatusCode: 500,
			Data:       false,
			Error:      nil,
		})
	}

	c.JSON(http.StatusOK, entities.Responce{
		StatusCode: 200,
		Data:       true,
		Error:      nil,
	})
}
