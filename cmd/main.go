package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"payment-service/config"
	"payment-service/internal/apiserver"
)

func main() {
	cfg := config.New()
	router := mux.NewRouter()
	log.Info("read configuration")
	if err := env.Parse(cfg); err != nil {
		log.Panic(err)
	}

	if err := apiserver.Start(cfg, router); err != nil {
		log.Fatal(err)
	}
}
