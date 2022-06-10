package handlers

import (
	"payment-service/internal/handlers"
	"payment-service/internal/services"
)

type Handlers struct {
	service             services.Service
	handlersTransaction *HandlerTransaction
}

func New(service services.Service) *Handlers {
	return &Handlers{
		service: service,
	}
}

func (h *Handlers) Transact() handlers.HandTransact {
	if h.handlersTransaction != nil {
		return h.handlersTransaction
	}

	h.handlersTransaction = &HandlerTransaction{
		handler: h,
	}
	return h.handlersTransaction
}
