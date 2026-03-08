package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rdl/core/internal/modules/payment/domain"
	"github.com/rdl/core/internal/modules/payment/service"
)

type PaymentHandler struct {
	svc service.PaymentService
}

func New(svc service.PaymentService) *PaymentHandler {
	return &PaymentHandler{svc: svc}
}

func (h *PaymentHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/payments", h.handleInitiate)
	mux.HandleFunc("GET /api/v1/payments", h.handleList)
	mux.HandleFunc("GET /api/v1/payments/{id}", h.handleGetByID)
	mux.HandleFunc("POST /api/v1/payments/webhook", h.handleWebhook)
}

func (h *PaymentHandler) handleInitiate(w http.ResponseWriter, r *http.Request) {
	var req domain.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := "user-placeholder"
	p, err := h.svc.InitiatePayment(r.Context(), userID, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (h *PaymentHandler) handleList(w http.ResponseWriter, r *http.Request) {
	filter := domain.ListPaymentsFilter{Page: 1, Limit: 20}
	payments, err := h.svc.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

func (h *PaymentHandler) handleGetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	p, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *PaymentHandler) handleWebhook(w http.ResponseWriter, r *http.Request) {
	var req domain.PaymentWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.svc.HandleWebhook(r.Context(), &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
