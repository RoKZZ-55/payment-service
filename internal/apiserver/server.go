package apiserver

import (
	"github.com/gorilla/mux"
	"net/http"
	"payment-service/internal/handlers/handlers"
	"payment-service/internal/middleware"
)

type Server struct {
	router     *mux.Router
	handler    *handlers.Handlers
	middleware *middleware.Auth
}

func newServer(router *mux.Router, handler *handlers.Handlers, middleware *middleware.Auth) *Server {
	s := &Server{
		router:     router,
		handler:    handler,
		middleware: middleware,
	}

	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/create-payment", s.handler.Payment().CreatePayment()).Methods("POST")
	s.router.HandleFunc("/change-payment-status", s.handler.Payment().ChangePaymentStatus()).Methods("PATCH")
	s.router.HandleFunc("/get-payment-status-by-id/{id}", s.handler.Payment().GetPaymentStatusByID()).Methods("GET")
	s.router.HandleFunc("/get-payments-by-userid/{id}", s.handler.Payment().GetPaymentsByID()).Methods("GET")
	s.router.HandleFunc("/get-payments-by-email/{email}", s.handler.Payment().GetPaymentsByEmail()).Methods("GET")
	s.router.HandleFunc("/cancel-payment-by-id/{id}", s.handler.Payment().CancelPaymentByID()).Methods("PATCH")
	s.router.Use(s.middleware.BasicAuth)
}
