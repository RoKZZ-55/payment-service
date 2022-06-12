package dbstorage

import (
	"database/sql"
	_ "github.com/lib/pq" // ...
	"payment-service/internal/storages"
)

type Storage struct {
	db             *sql.DB
	paymentStorage *PaymentStorage
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Payment() storages.PaymentStorage {
	if s.paymentStorage != nil {
		return s.paymentStorage
	}

	s.paymentStorage = &PaymentStorage{
		storage: s,
	}
	return s.paymentStorage
}
