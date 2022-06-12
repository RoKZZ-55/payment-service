package handlers

import (
	"payment-service/internal/handlers"
	"payment-service/internal/services"
)

type Handlers struct {
	service        services.Service
	paymentHandler *PaymentHandler
}

func New(service services.Service) *Handlers {
	return &Handlers{
		service: service,
	}
}

func (h *Handlers) Payment() handlers.PaymentHandler {
	if h.paymentHandler != nil {
		return h.paymentHandler
	}

	h.paymentHandler = &PaymentHandler{
		handler: h,
	}
	return h.paymentHandler
}
