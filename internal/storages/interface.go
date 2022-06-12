package storages

import "payment-service/internal/model"

type Storage interface {
	Payment() PaymentStorage
}

type PaymentStorage interface {
	CreatePayment(transact *model.Transactions) error
	ChangePaymentStatus(transact *model.Transactions) error
	GetPaymentStatusByID(transactID uint64) (status string, err error)
	GetPaymentsByID(userID uint64) (transact []model.Transactions, err error)
	GetPaymentsByEmail(email string) (transact []model.Transactions, err error)
	CancelPaymentByID(transactID uint64) error
}
