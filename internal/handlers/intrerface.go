package handlers

import "net/http"

type Handlers interface {
	Payment() PaymentHandler
}

type PaymentHandler interface {
	Create() http.HandlerFunc
	ChangeStatus() http.HandlerFunc
	GetStatus() http.HandlerFunc
	GetByEmailOrID() http.HandlerFunc
	Cancel() http.HandlerFunc
}
