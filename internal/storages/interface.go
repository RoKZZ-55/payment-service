package storages

import "payment-service/internal/model"

type Storage interface {
	Transactions() TransactionsStorage
}

type TransactionsStorage interface {
	CreatePayment(transact *model.Transactions) error
}
