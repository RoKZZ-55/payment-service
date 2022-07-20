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

func (h *PaymentHandler) Create() http.HandlerFunc {
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
		t := &model.Transaction{
			UserID:   req.ID,
			Email:    req.Email,
			Sum:      req.Sum,
			Currency: req.Currency,
		}

		if err := h.handler.service.Payment().Create(t); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusCreated, "Payment created")
	}
}

func (h *PaymentHandler) ChangeStatus() http.HandlerFunc {
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
		t := &model.Transaction{
			TransactID: req.TransactID,
			Status:     req.Status,
		}
		if err := h.handler.service.Payment().ChangeStatus(t); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, "Payment status changed successfully")
	}
}

func (h *PaymentHandler) GetStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		data, err := h.handler.service.Payment().GetStatus(id)
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *PaymentHandler) GetByEmailOrID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		path := vars["path"]
		data, err := h.handler.service.Payment().GetByEmailOrID(path)
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *PaymentHandler) Cancel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		if err := h.handler.service.Payment().Cancel(id); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, "Payment canceled successfully")
	}
}
