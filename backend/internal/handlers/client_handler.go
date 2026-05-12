package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/leonardosth/onboardly/internal/models"
	"github.com/leonardosth/onboardly/internal/service"
)

type ClientHandler struct {
	service *service.ClientService
}

func NewClientHandler(s *service.ClientService) *ClientHandler {
	return &ClientHandler{service: s}
}

func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	if c.Nome == "" || c.CNPJ == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Nome e CNPJ são obrigatórios"})
		return
	}

	err := h.service.CreateClient(r.Context(), &c)
	if err != nil {
		if err.Error() == "cliente com esse CNPJ já cadastrado" {
			respondJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao criar cliente", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao salvar cliente"})
		return
	}

	respondJSON(w, http.StatusCreated, c)
}

func (h *ClientHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.service.GetEveryone(r.Context())
	if err != nil {
		slog.Error("Erro ao buscar clientes", "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar clientes"})
		return
	}

	if clientes == nil {
		clientes = []*models.Cliente{}
	}

	respondJSON(w, http.StatusOK, clientes)
}

func (h *ClientHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/clientes/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	cliente, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		slog.Error("Erro ao buscar cliente por ID", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar cliente"})
		return
	}

	if cliente == nil {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "Cliente não encontrado"})
		return
	}

	respondJSON(w, http.StatusOK, cliente)
}

func (h *ClientHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/clientes/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	c.ID = id
	err = h.service.Update(r.Context(), &c)
	if err != nil {
		if err.Error() == "cliente não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao atualizar cliente", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao atualizar cliente"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Cliente atualizado com sucesso"})
}

func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r, "/clientes/")
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "cliente não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		slog.Error("Erro ao deletar cliente", "id", id, "error", err)
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao deletar cliente"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Cliente deletado com sucesso"})
}
