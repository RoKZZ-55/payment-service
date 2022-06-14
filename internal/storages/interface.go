package storages

import "payment-service/internal/model"

type Storage interface {
	Payment() PaymentStorage
}

type PaymentStorage interface {
	CreatePayment(transact *model.Transaction) error
	ChangePaymentStatus(transact *model.Transaction) error
	GetPaymentStatusByID(transactID uint64) (status string, err error)
	GetPaymentsByID(userID uint64) (transact []model.Transaction, err error)
	GetPaymentsByEmail(email string) (transact []model.Transaction, err error)
	CancelPaymentByID(transactID uint64) error
}
