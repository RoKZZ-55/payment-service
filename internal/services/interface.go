package services

import "payment-service/internal/model"

type Service interface {
	Transact() TransactionsService
}

type TransactionsService interface {
	CreatePayment(transact *model.Transactions) error
}
