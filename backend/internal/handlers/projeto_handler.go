package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/leonardosth/onboardly/internal/models"
	"github.com/leonardosth/onboardly/internal/service"
)

type ProjetoHandler struct {
	service *service.ProjetoService
}

func NewProjetoHandler(s *service.ProjetoService) *ProjetoHandler {
	return &ProjetoHandler{service: s}
}

func (h *ProjetoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Projeto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	err := h.service.Create(r.Context(), &p)
	if err != nil {
		slog.Error("Erro ao criar projeto", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao salvar projeto"})
		return
	}

	respondJSON(w, http.StatusCreated, p)
}

func (h *ProjetoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	projetos, err := h.service.GetEveryone(r.Context())
	if err != nil {
		slog.Error("Erro ao buscar projetos", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar projetos"})
		return
	}

	if projetos == nil {
		projetos = []*models.Projeto{}
	}

	respondJSON(w, http.StatusOK, projetos)
}

func (h *ProjetoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/projetos/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	projeto, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		slog.Error("Erro ao buscar projeto por ID", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar projeto"})
		return
	}

	if projeto == nil {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "Projeto não encontrado"})
		return
	}

	respondJSON(w, http.StatusOK, projeto)
}

func (h *ProjetoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/projetos/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	var p models.Projeto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	p.ID = id
	err = h.service.Update(r.Context(), &p)
	if err != nil {
		if err.Error() == "projeto não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao atualizar projeto", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao atualizar projeto"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Projeto atualizado com sucesso"})
}

func (h *ProjetoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/projetos/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "projeto não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao deletar projeto", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao deletar projeto"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Projeto deletado com sucesso"})
}

func (h *ProjetoHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.service.GetStats(r.Context())
	if err != nil {
		slog.Error("Erro ao buscar estatísticas do dashboard", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar estatísticas"})
		return
	}

	respondJSON(w, http.StatusOK, stats)
}
