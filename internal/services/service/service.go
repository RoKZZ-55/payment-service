package service

import (
	"payment-service/internal/services"
	"payment-service/internal/storages"
)

type Service struct {
	storage        storages.Storage
	paymentService *PaymentService
}

func New(storage storages.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Payment() services.PaymentService {
	if s.paymentService != nil {
		return s.paymentService
	}

	s.paymentService = &PaymentService{
		service: s,
	}
	return s.paymentService
}
