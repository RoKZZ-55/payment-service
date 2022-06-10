package service

import "payment-service/internal/model"

type TransactionsService struct {
	service *Service
}

func (s *TransactionsService) CreatePayment(t *model.Transactions) error {
	return s.service.storage.Transactions().CreatePayment(t)
}
