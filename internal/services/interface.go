package services

import "payment-service/internal/model"

type Service interface {
	Payment() PaymentService
}

type PaymentService interface {
	Create(transact *model.Transaction) error
	ChangeStatus(transact *model.Transaction) error
	GetStatus(transactID uint64) (status string, err error)
	GetByEmailOrID(path string) (transact []model.Transaction, err error)
	Cancel(transactID uint64) error
}
