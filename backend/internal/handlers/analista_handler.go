package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/leonardosth/onboardly/internal/models"
	"github.com/leonardosth/onboardly/internal/service"
)

type AnalistaHandler struct {
	service *service.AnalistaService
}

func NewAnalistaHandler(s *service.AnalistaService) *AnalistaHandler {
	return &AnalistaHandler{service: s}
}

func (h *AnalistaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var a models.Analista
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	if a.Nome == "" || a.Email == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Nome e Email são obrigatórios"})
		return
	}

	err := h.service.Create(r.Context(), &a)
	if err != nil {
		if err.Error() == "analista com esse email já cadastrado" {
			respondJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao criar analista", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao salvar analista"})
		return
	}

	respondJSON(w, http.StatusCreated, a)
}

func (h *AnalistaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	analistas, err := h.service.GetEveryone(r.Context())
	if err != nil {
		slog.Error("Erro ao buscar analistas", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar analistas"})
		return
	}

	if analistas == nil {
		analistas = []*models.Analista{}
	}

	respondJSON(w, http.StatusOK, analistas)
}

func (h *AnalistaHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/analistas/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	analista, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		slog.Error("Erro ao buscar analista por ID", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar analista"})
		return
	}

	if analista == nil {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "Analista não encontrado"})
		return
	}

	respondJSON(w, http.StatusOK, analista)
}

func (h *AnalistaHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/analistas/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	var a models.Analista
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	a.ID = id
	err = h.service.Update(r.Context(), &a)
	if err != nil {
		if err.Error() == "analista não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao atualizar analista", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao atualizar analista"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Analista atualizado com sucesso"})
}

func (h *AnalistaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/analistas/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "analista não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao deletar analista", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao deletar analista"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Analista deletado com sucesso"})
}
