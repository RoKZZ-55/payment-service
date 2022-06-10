package handlers

import "net/http"

type Hand interface {
	Transact() HandTransact
}

type HandTransact interface {
	HandCreatePayment() http.HandlerFunc
}
