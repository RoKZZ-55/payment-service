package handlers

import "net/http"

type Handlers interface {
	Payment() PaymentHandler
}

type PaymentHandler interface {
	CreatePayment() http.HandlerFunc
	ChangePaymentStatus() http.HandlerFunc
	GetPaymentStatusByID() http.HandlerFunc
	GetPaymentsByID() http.HandlerFunc
	GetPaymentsByEmail() http.HandlerFunc
	CancelPaymentByID() http.HandlerFunc
}
