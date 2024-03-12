package entities

type Responce struct {
	StatusCode int         `json:"StatusCode,omitempty"`
	Message    string      `json:"Message,omitempty"`
	Error      error       `json:"Error,omitempty"`
	Data       interface{} `json:"Data,omitempty"`
}

type MakePaymentUsecase struct {
	OrderID        string
	UserID         string
	IsPayed        bool
	Price          uint
}