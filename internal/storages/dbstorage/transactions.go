package dbstorage

import "payment-service/internal/model"

type TransactionsStorage struct {
	storage *Storage
}

func (s *TransactionsStorage) CreatePayment(t *model.Transactions) error {
	return nil
}
