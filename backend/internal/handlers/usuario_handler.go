package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/leonardosth/onboardly/internal/models"
	"github.com/leonardosth/onboardly/internal/service"
)

type UsuarioHandler struct {
	service *service.UsuarioService
}

func NewUsuarioHandler(s *service.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{service: s}
}

func (h *UsuarioHandler) CreateAnalista(w http.ResponseWriter, r *http.Request) {
	var u models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	u.Cargo = "Analista" // Força o cargo como Analista neste endpoint

	if !validatePayload(w, u) {
		return
	}

	err := h.service.Create(r.Context(), &u)
	if err != nil {
		if err.Error() == "usuário com esse email já cadastrado" {
			respondJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao criar analista", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao salvar analista"})
		return
	}

	respondJSON(w, http.StatusCreated, u)
}

func (h *UsuarioHandler) GetAnalistas(w http.ResponseWriter, r *http.Request) {
	analistas, err := h.service.GetByCargo(r.Context(), "Analista")
	if err != nil {
		slog.Error("Erro ao buscar analistas", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar analistas"})
		return
	}

	if analistas == nil {
		analistas = []*models.Usuario{}
	}

	respondJSON(w, http.StatusOK, analistas)
}

func (h *UsuarioHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/analistas/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	usuario, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		slog.Error("Erro ao buscar usuário por ID", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar usuário"})
		return
	}

	if usuario == nil {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "Usuário não encontrado"})
		return
	}

	respondJSON(w, http.StatusOK, usuario)
}

func (h *UsuarioHandler) UpdateAnalista(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/analistas/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	var u models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	u.ID = id
	u.Cargo = "Analista"

	if !validatePayload(w, u) {
		return
	}

	err = h.service.Update(r.Context(), &u)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao atualizar analista", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao atualizar analista"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Analista atualizado com sucesso"})
}

func (h *UsuarioHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/analistas/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao deletar usuário", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao deletar usuário"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Analista deletado com sucesso"})
}
