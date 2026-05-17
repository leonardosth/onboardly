package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/leonardosth/onboardly/internal/models"
	"github.com/leonardosth/onboardly/internal/service"
)

type ReuniaoHandler struct {
	service *service.ReuniaoService
}

func NewReuniaoHandler(s *service.ReuniaoService) *ReuniaoHandler {
	return &ReuniaoHandler{service: s}
}

func (h *ReuniaoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var re models.Reuniao
	if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	// Se analista_id não foi enviado, tenta pegar do token JWT
	if re.AnalistaID == uuid.Nil {
		val := r.Context().Value(UserContextKey)
		if val != nil {
			if claims, ok := val.(*jwt.MapClaims); ok {
				if sub, ok := (*claims)["sub"].(string); ok {
					if id, err := uuid.Parse(sub); err == nil {
						re.AnalistaID = id
					}
				}
			}
		}
	}

	err := h.service.Create(r.Context(), &re)
	if err != nil {
		slog.Error("Erro ao criar reunião", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao salvar reunião"})
		return
	}

	respondJSON(w, http.StatusCreated, re)
}

func (h *ReuniaoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	reunioes, err := h.service.GetEveryone(r.Context())
	if err != nil {
		slog.Error("Erro ao buscar reuniões", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar reuniões"})
		return
	}

	if reunioes == nil {
		reunioes = []*models.Reuniao{}
	}

	respondJSON(w, http.StatusOK, reunioes)
}

func (h *ReuniaoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/reunioes/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	reuniao, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		slog.Error("Erro ao buscar reunião por ID", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar reunião"})
		return
	}

	if reuniao == nil {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "Reunião não encontrada"})
		return
	}

	respondJSON(w, http.StatusOK, reuniao)
}

func (h *ReuniaoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/reunioes/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	var re models.Reuniao
	if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	re.ID = id
	err = h.service.Update(r.Context(), &re)
	if err != nil {
		if err.Error() == "reunião não encontrada" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao atualizar reunião", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao atualizar reunião"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Reunião atualizada com sucesso"})
}

func (h *ReuniaoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/reunioes/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "reunião não encontrada" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao deletar reunião", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao deletar reunião"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Reunião deletada com sucesso"})
}
