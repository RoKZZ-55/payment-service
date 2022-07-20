package storages

import "payment-service/internal/model"

type Storage interface {
	Payment() PaymentStorage
}

type PaymentStorage interface {
	Create(transact *model.Transaction) error
	ChangeStatus(transact *model.Transaction) error
	GetStatus(transactID uint64) (status string, err error)
	GetByEmailOrID(path *model.Transaction) (transact []model.Transaction, err error)
	Cancel(transactID uint64) error
}
