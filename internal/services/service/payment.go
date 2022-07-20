package service

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	"payment-service/internal/model"
	"strconv"
)

type PaymentService struct {
	service *Service
}

func (s *PaymentService) Create(t *model.Transaction) error {
	if !valid.IsEmail(t.Email) {
		return errors.New("invalid email format")
	}
	if len(t.Currency) != 3 {
		return errors.New("invalid currency format, currency name consists of 3 characters")
	}
	return s.service.storage.Payment().Create(t)
}

func (s *PaymentService) ChangeStatus(t *model.Transaction) error {
	switch t.Status {
	case "УСПЕХ", "НЕУСПЕХ", "ОШИБКА":
		return s.service.storage.Payment().ChangeStatus(t)
	default:
		return errors.New("wrong status")
	}
}

func (s *PaymentService) GetStatus(transactID uint64) (status string, err error) {
	return s.service.storage.Payment().GetStatus(transactID)
}

func (s *PaymentService) GetByEmailOrID(path string) (transact []model.Transaction, err error) {
	p := new(model.Transaction)
	if valid.IsEmail(path) {
		p.Email = path
		return s.service.storage.Payment().GetByEmailOrID(p)
	}
	if p.UserID, err = strconv.ParseUint(path, 10, 64); err == nil {
		return s.service.storage.Payment().GetByEmailOrID(p)
	}
	return nil, errors.New("incorrect path, you must specify an email or user id")
}

func (s *PaymentService) Cancel(transactID uint64) error {
	return s.service.storage.Payment().Cancel(transactID)
}
