package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rdl/core/internal/modules/project/domain"
	"github.com/rdl/core/internal/modules/project/service"
)

type ProjectHandler struct {
	svc service.ProjectService
}

func New(svc service.ProjectService) *ProjectHandler {
	return &ProjectHandler{svc: svc}
}

func (h *ProjectHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/projects", h.handleCreate)
	mux.HandleFunc("GET /api/v1/projects", h.handleList)
	mux.HandleFunc("GET /api/v1/projects/{id}", h.handleGetByID)
	mux.HandleFunc("PUT /api/v1/projects/{id}", h.handleUpdate)
	mux.HandleFunc("DELETE /api/v1/projects/{id}", h.handleDelete)
}

func (h *ProjectHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: extract ownerID from context
	ownerID := "owner-placeholder"
	p, err := h.svc.Create(r.Context(), ownerID, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (h *ProjectHandler) handleList(w http.ResponseWriter, r *http.Request) {
	filter := domain.ListProjectsFilter{Page: 1, Limit: 20}
	projects, err := h.svc.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) handleGetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	p, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *ProjectHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req domain.UpdateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := h.svc.Update(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *ProjectHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
