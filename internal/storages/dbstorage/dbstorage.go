package dbstorage

import (
	"database/sql"
	_ "github.com/lib/pq" // ...
	"payment-service/internal/storages"
)

type Storage struct {
	db                  *sql.DB
	transactionsStorage *TransactionsStorage
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Transactions() storages.TransactionsStorage {
	if s.transactionsStorage != nil {
		return s.transactionsStorage
	}

	s.transactionsStorage = &TransactionsStorage{
		storage: s,
	}
	return s.transactionsStorage
}
