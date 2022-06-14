package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"payment-service/internal/model"
	"payment-service/internal/utils"
	"strconv"
)

type PaymentHandler struct {
	handler *Handlers
}

func (h *PaymentHandler) CreatePayment() http.HandlerFunc {
	type request struct {
		ID       uint64  `json:"user-id"`
		Email    string  `json:"email"`
		Sum      float64 `json:"sum"`
		Currency string  `json:"currency"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		t := &model.Transactions{
			UserID:   req.ID,
			Email:    req.Email,
			Sum:      req.Sum,
			Currency: req.Currency,
		}

		if err := h.handler.service.Payment().CreatePayment(t); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusCreated, "Payment created")
	}
}

func (h *PaymentHandler) ChangePaymentStatus() http.HandlerFunc {
	type request struct {
		TransactID uint64 `json:"transact-id"`
		Status     string `json:"status"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		t := &model.Transactions{
			TransactID: req.TransactID,
			Status:     req.Status,
		}
		if err := h.handler.service.Payment().ChangePaymentStatus(t); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusCreated, "Payment status changed successfully")
	}
}

func (h *PaymentHandler) GetPaymentStatusByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		data, err := h.handler.service.Payment().GetPaymentStatusByID(id)
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *PaymentHandler) GetPaymentsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		data, err := h.handler.service.Payment().GetPaymentsByID(id)
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *PaymentHandler) GetPaymentsByEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email := vars["email"]
		data, err := h.handler.service.Payment().GetPaymentsByEmail(email)
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *PaymentHandler) CancelPaymentByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		if err := h.handler.service.Payment().CancelPaymentByID(id); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, "Payment canceled successfully")
	}
}
