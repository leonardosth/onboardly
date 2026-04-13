package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/leonardosth/onboardly/internal/models"
	"github.com/leonardosth/onboardly/internal/service"
)

type ClientHandler struct {
	service *service.ClientService
}

func NewClientHandler(s *service.ClientService) *ClientHandler {
	return &ClientHandler{service: s}
}

// JSON auxiliar para simplificar respostas
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Payload inválido"})
		return
	}

	// Validações básicas omitidas para simplificar, mas recomendadas
	if c.NomeFantasia == "" || c.CNPJ == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Nome e CNPJ são obrigatórios"})
		return
	}

	err := h.service.CreateClient(r.Context(), &c)
	if err != nil {
		if err.Error() == "cliente com esse CNPJ já cadastrado" {
			respondJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
			return
		}
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao salvar cliente"})
		return
	}

	respondJSON(w, http.StatusCreated, c)
}

func (h *ClientHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.service.GetEveryone(r.Context())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao buscar clientes"})
		return
	}

	// Garantir que um slice nulo retorne [] no JSON ao invés de null
	if clientes == nil {
		clientes = []*models.Cliente{}
	}

	respondJSON(w, http.StatusOK, clientes)
}

func (h *ClientHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	cliente, err := h.service.GetByID(r.Context(), id)
	if err != nil {
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
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
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
	err = h.service.UpdateCliente(r.Context(), &c)
	if err != nil {
		if err.Error() == "cliente não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao atualizar cliente"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Cliente atualizado com sucesso"})
}

func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return
	}

	err = h.service.DeleteCliente(r.Context(), id)
	if err != nil {
		if err.Error() == "cliente não encontrado" {
			respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno ao deletar cliente"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Cliente deletado com sucesso"})
}
