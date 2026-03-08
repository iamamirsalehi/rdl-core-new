package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rdl/core/internal/modules/subscription/domain"
	"github.com/rdl/core/internal/modules/subscription/service"
)

type SubscriptionHandler struct {
	svc service.SubscriptionService
}

func New(svc service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{svc: svc}
}

func (h *SubscriptionHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/plans", h.handleListPlans)
	mux.HandleFunc("POST /api/v1/subscriptions", h.handleSubscribe)
	mux.HandleFunc("GET /api/v1/subscriptions", h.handleList)
	mux.HandleFunc("GET /api/v1/subscriptions/{id}", h.handleGetByID)
	mux.HandleFunc("DELETE /api/v1/subscriptions/{id}", h.handleCancel)
}

func (h *SubscriptionHandler) handleListPlans(w http.ResponseWriter, r *http.Request) {
	plans, err := h.svc.ListPlans(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plans)
}

func (h *SubscriptionHandler) handleSubscribe(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := "user-placeholder"
	sub, err := h.svc.Subscribe(r.Context(), userID, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sub)
}

func (h *SubscriptionHandler) handleList(w http.ResponseWriter, r *http.Request) {
	filter := domain.ListSubscriptionsFilter{Page: 1, Limit: 20}
	subs, err := h.svc.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subs)
}

func (h *SubscriptionHandler) handleGetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	sub, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)
}

func (h *SubscriptionHandler) handleCancel(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.svc.Cancel(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
