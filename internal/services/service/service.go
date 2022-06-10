package service

import (
	"payment-service/internal/services"
	"payment-service/internal/storages"
)

type Service struct {
	storage             storages.Storage
	transactionsService *TransactionsService
}

func New(storage storages.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Transact() services.TransactionsService {
	if s.transactionsService != nil {
		return s.transactionsService
	}

	s.transactionsService = &TransactionsService{
		service: s,
	}
	return s.transactionsService
}
