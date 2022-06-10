package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"payment-service/config"
	"payment-service/internal/handlers/handlers"
	"payment-service/internal/services/service"
	"payment-service/internal/storages/dbstorage"
)

func Start(cfg *config.Config, router *mux.Router) error {
	log.Info("connecting to postgres")
	db, err := newDB(cfg.Postgres)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	log.Info("payment service launch")
	srv := newServer(router, handlers.New(service.New(dbstorage.New(db))))
	return http.ListenAndServe(cfg.BindAddr, srv)
}

func newDB(pq config.Postgres) (*sql.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", pq.User, pq.Password, pq.Host, pq.DBName, pq.SSL)
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
