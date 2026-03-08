package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rdl/core/internal/modules/evaluation/domain"
	"github.com/rdl/core/internal/modules/evaluation/service"
)

type EvaluationHandler struct {
	svc service.EvaluationService
}

func New(svc service.EvaluationService) *EvaluationHandler {
	return &EvaluationHandler{svc: svc}
}

func (h *EvaluationHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/evaluations", h.handleCreate)
	mux.HandleFunc("GET /api/v1/evaluations", h.handleList)
	mux.HandleFunc("GET /api/v1/evaluations/{id}", h.handleGetByID)
	mux.HandleFunc("PUT /api/v1/evaluations/{id}", h.handleUpdate)
	mux.HandleFunc("DELETE /api/v1/evaluations/{id}", h.handleDelete)
}

func (h *EvaluationHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateEvaluationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	evaluatorID := "evaluator-placeholder"
	e, err := h.svc.Create(r.Context(), evaluatorID, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func (h *EvaluationHandler) handleList(w http.ResponseWriter, r *http.Request) {
	filter := domain.ListEvaluationsFilter{Page: 1, Limit: 20}
	items, err := h.svc.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *EvaluationHandler) handleGetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	e, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func (h *EvaluationHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req domain.UpdateEvaluationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e, err := h.svc.Update(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func (h *EvaluationHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
