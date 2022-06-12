package service

import (
	"errors"
	"payment-service/internal/model"
)

type PaymentService struct {
	service *Service
}

func (s *PaymentService) CreatePayment(t *model.Transactions) error {
	return s.service.storage.Payment().CreatePayment(t)
}

func (s *PaymentService) ChangePaymentStatus(t *model.Transactions) error {
	switch t.Status {
	case "УСПЕХ", "НЕУСПЕХ", "ОШИБКА":
		return s.service.storage.Payment().ChangePaymentStatus(t)
	default:
		return errors.New("wrong status")
	}
}

func (s *PaymentService) GetPaymentStatusByID(transactID uint64) (status string, err error) {
	return s.service.storage.Payment().GetPaymentStatusByID(transactID)
}

func (s *PaymentService) GetPaymentsByID(userID uint64) (transact []model.Transactions, err error) {
	return s.service.storage.Payment().GetPaymentsByID(userID)
}

func (s *PaymentService) GetPaymentsByEmail(email string) (transact []model.Transactions, err error) {
	return s.service.storage.Payment().GetPaymentsByEmail(email)
}

func (s *PaymentService) CancelPaymentByID(transactID uint64) error {
	return s.service.storage.Payment().CancelPaymentByID(transactID)
}
