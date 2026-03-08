package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rdl/core/internal/modules/cathub/domain"
	"github.com/rdl/core/internal/modules/cathub/service"
)

type CathubHandler struct {
	svc service.CathubService
}

func New(svc service.CathubService) *CathubHandler {
	return &CathubHandler{svc: svc}
}

func (h *CathubHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/hubs", h.handleCreateHub)
	mux.HandleFunc("GET /api/v1/hubs", h.handleListHubs)
	mux.HandleFunc("GET /api/v1/hubs/{id}", h.handleGetHub)
	mux.HandleFunc("PUT /api/v1/hubs/{id}", h.handleUpdateHub)
	mux.HandleFunc("DELETE /api/v1/hubs/{id}", h.handleDeleteHub)
}

func (h *CathubHandler) handleCreateHub(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateHubRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ownerID := "owner-placeholder"
	hub, err := h.svc.CreateHub(r.Context(), ownerID, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hub)
}

func (h *CathubHandler) handleListHubs(w http.ResponseWriter, r *http.Request) {
	filter := domain.ListHubsFilter{Page: 1, Limit: 20}
	hubs, err := h.svc.ListHubs(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hubs)
}

func (h *CathubHandler) handleGetHub(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	hub, err := h.svc.GetHubByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hub)
}

func (h *CathubHandler) handleUpdateHub(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req domain.UpdateHubRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hub, err := h.svc.UpdateHub(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hub)
}

func (h *CathubHandler) handleDeleteHub(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.svc.DeleteHub(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
