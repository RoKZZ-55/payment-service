package service

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	"payment-service/internal/model"
)

type PaymentService struct {
	service *Service
}

func (s *PaymentService) CreatePayment(t *model.Transaction) error {
	if !valid.IsEmail(t.Email) {
		return errors.New("invalid email format")
	}
	if len(t.Currency) != 3 {
		return errors.New("invalid currency format, currency name consists of 3 characters")
	}
	return s.service.storage.Payment().CreatePayment(t)
}

func (s *PaymentService) ChangePaymentStatus(t *model.Transaction) error {
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

func (s *PaymentService) GetPaymentsByID(userID uint64) (transact []model.Transaction, err error) {
	return s.service.storage.Payment().GetPaymentsByID(userID)
}

func (s *PaymentService) GetPaymentsByEmail(email string) (transact []model.Transaction, err error) {
	if !valid.IsEmail(email) {
		return nil, errors.New("invalid email format")
	}
	return s.service.storage.Payment().GetPaymentsByEmail(email)
}

func (s *PaymentService) CancelPaymentByID(transactID uint64) error {
	return s.service.storage.Payment().CancelPaymentByID(transactID)
}
