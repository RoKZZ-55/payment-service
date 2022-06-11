package dbstorage

import (
	"payment-service/internal/model"
)

type TransactionsStorage struct {
	storage *Storage
}

func (s *TransactionsStorage) CreatePayment(t *model.Transactions) error {
	return s.storage.db.QueryRow(
		"INSERT INTO transact (user_id, email, sum, currency, date_time_create, date_time_last_change, status) "+
			"VALUES ($1, $2, $3, $4, now(), now(), $5)",
		t.UserID,
		t.Email,
		t.Sum,
		t.Currency,
		"НОВЫЙ",
	).Err()
}
