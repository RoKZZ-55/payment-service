package services

import "payment-service/internal/model"

type Service interface {
	Payment() PaymentService
}

type PaymentService interface {
	CreatePayment(transact *model.Transaction) error
	ChangePaymentStatus(transact *model.Transaction) error
	GetPaymentStatusByID(transactID uint64) (status string, err error)
	GetPaymentsByID(userID uint64) (transact []model.Transaction, err error)
	GetPaymentsByEmail(email string) (transact []model.Transaction, err error)
	CancelPaymentByID(transactID uint64) error
}
