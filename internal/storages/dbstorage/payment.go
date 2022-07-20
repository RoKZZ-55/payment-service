package dbstorage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"payment-service/internal/model"
)

type PaymentStorage struct {
	storage *Storage
}

func (s *PaymentStorage) Create(t *model.Transaction) error {
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

func (s *PaymentStorage) ChangeStatus(t *model.Transaction) error {
	var statusDB string
	err := s.storage.db.QueryRow("SELECT status FROM transact WHERE transact_id = $1",
		t.TransactID,
	).Scan(&statusDB)
	if err != nil {
		return err
	}
	if statusDB == "УСПЕХ" || statusDB == "НЕУСПЕХ" {
		return errors.New(fmt.Sprintf("statuses 'УСПЕХ', 'НЕУСПЕХ' cannot be changed "))
	}
	err = s.storage.db.QueryRow(
		"UPDATE transact SET status = $1, date_time_last_change = now() WHERE transact_id = $2",
		t.Status,
		t.TransactID,
	).Err()
	return err
}

func (s *PaymentStorage) GetStatus(transactID uint64) (status string, err error) {
	err = s.storage.db.QueryRow("SELECT status FROM transact WHERE transact_id = $1", transactID).Scan(&status)
	return status, err
}

func (s *PaymentStorage) GetByEmailOrID(p *model.Transaction) (transact []model.Transaction, err error) {
	var rows *sql.Rows
	if p.Email != "" && p.UserID == 0 {
		rows, err = s.storage.db.Query(
			"SELECT * FROM transact WHERE email = $1 ORDER BY date_time_last_change DESC",
			p.Email,
		)
	} else if p.UserID != 0 && p.Email == "" {
		rows, err = s.storage.db.Query(
			"SELECT * FROM transact WHERE user_id = $1 ORDER BY date_time_last_change DESC",
			p.UserID,
		)
	} else {
		return nil, errors.New("internal error")
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := model.Transaction{}
		err := rows.Scan(
			&t.TransactID, &t.UserID, &t.Email, &t.Sum, &t.Currency,
			&t.DateTimeCreate, &t.DateTimeLastChange, &t.Status,
		)
		if err != nil {
			logrus.Error(err)
			continue
		}
		transact = append(transact, t)
	}
	if transact == nil {
		return nil, errors.New("transactions not found")
	}
	return transact, err
}

func (s *PaymentStorage) Cancel(transactID uint64) error {
	var statusDB string
	err := s.storage.db.QueryRow("SELECT status FROM transact WHERE transact_id = $1", transactID).Scan(&statusDB)
	if err != nil {
		return err
	}
	if statusDB == "УСПЕХ" || statusDB == "НЕУСПЕХ" {
		return errors.New(fmt.Sprintf("statuses 'УСПЕХ', 'НЕУСПЕХ' cannot be changed "))
	}
	err = s.storage.db.QueryRow(
		"UPDATE transact SET status = $1, date_time_last_change = now() WHERE transact_id = $2",
		"ОТМЕНЕН",
		transactID,
	).Err()
	return err
}
