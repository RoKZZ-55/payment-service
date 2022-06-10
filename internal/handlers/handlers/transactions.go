package handlers

import (
	"encoding/json"
	"net/http"
	"payment-service/internal/model"
	"payment-service/internal/utils"
)

type HandlerTransaction struct {
	handler *Handlers
}

func (h *HandlerTransaction) HandCreatePayment() http.HandlerFunc {
	type request struct {
		ID       uint    `json:"id"`
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

		u := &model.User{
			ID:    req.ID,
			Email: req.Email,
		}
		t := &model.Transactions{
			User:     *u,
			Sum:      req.Sum,
			Currency: req.Currency,
		}

		if err := h.handler.service.Transact().CreatePayment(t); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
	}
}
