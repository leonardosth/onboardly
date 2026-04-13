package service

import (
	"context"
	"testing"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

// Mock do Repositório
type mockClientRepository struct {
	mockGetByCNPJ   func(cnpj string) (*models.Cliente, error)
	mockGetByID     func(id uuid.UUID) (*models.Cliente, error)
	mockCreate      func(cliente *models.Cliente) error
	mockGetEveryone func() ([]*models.Cliente, error)
	mockUpdate      func(cliente *models.Cliente) error
	mockDelete      func(id uuid.UUID) error
}

func (m *mockClientRepository) GetByCNPJ(ctx context.Context, cnpj string) (*models.Cliente, error) {
	return m.mockGetByCNPJ(cnpj)
}
func (m *mockClientRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Cliente, error) {
	return m.mockGetByID(id)
}
func (m *mockClientRepository) Create(ctx context.Context, cliente *models.Cliente) error {
	return m.mockCreate(cliente)
}
func (m *mockClientRepository) GetEveryone(ctx context.Context) ([]*models.Cliente, error) {
	return m.mockGetEveryone()
}
func (m *mockClientRepository) UpdateCliente(ctx context.Context, cliente *models.Cliente) error {
	return m.mockUpdate(cliente)
}
func (m *mockClientRepository) DeleteCliente(ctx context.Context, id uuid.UUID) error {
	return m.mockDelete(id)
}

func TestCreateClient(t *testing.T) {
	t.Run("Deve retornar erro se CNPJ ja existir", func(t *testing.T) {
		mockRepo := &mockClientRepository{
			mockGetByCNPJ: func(cnpj string) (*models.Cliente, error) {
				return &models.Cliente{CNPJ: "123"}, nil // Simulando cliente já existente
			},
		}
		service := NewClientService(mockRepo)

		err := service.CreateClient(context.Background(), &models.Cliente{CNPJ: "123"})
		if err == nil || err.Error() != "cliente com esse CNPJ já cadastrado" {
			t.Errorf("Esperava erro de CNPJ duplicado, recebeu: %v", err)
		}
	})

	t.Run("Deve criar cliente se CNPJ for novo", func(t *testing.T) {
		mockRepo := &mockClientRepository{
			mockGetByCNPJ: func(cnpj string) (*models.Cliente, error) {
				return nil, nil // Simulando cliente não encontrado
			},
			mockCreate: func(cliente *models.Cliente) error {
				return nil
			},
		}
		service := NewClientService(mockRepo)

		err := service.CreateClient(context.Background(), &models.Cliente{CNPJ: "456"})
		if err != nil {
			t.Errorf("Nao esperava erro, recebeu: %v", err)
		}
	})
}

func TestGetEveryone(t *testing.T) {
	t.Run("Deve retornar lista de clientes", func(t *testing.T) {
		expectedClients := []*models.Cliente{
			{CNPJ: "123", NomeFantasia: "Cliente 1"},
			{CNPJ: "456", NomeFantasia: "Cliente 2"},
		}
		mockRepo := &mockClientRepository{
			mockGetEveryone: func() ([]*models.Cliente, error) {
				return expectedClients, nil
			},
		}
		service := NewClientService(mockRepo)

		clients, err := service.GetEveryone(context.Background())
		if err != nil {
			t.Errorf("Não esperava erro, recebeu: %v", err)
		}
		if len(clients) != 2 {
			t.Errorf("Esperava 2 clientes, recebeu: %d", len(clients))
		}
	})
}

func TestGetByCNPJ(t *testing.T) {
	t.Run("Deve retornar cliente pelo CNPJ", func(t *testing.T) {
		expectedClient := &models.Cliente{CNPJ: "123", NomeFantasia: "Cliente 1"}
		mockRepo := &mockClientRepository{
			mockGetByCNPJ: func(cnpj string) (*models.Cliente, error) {
				return expectedClient, nil
			},
		}
		service := NewClientService(mockRepo)

		client, err := service.GetByCNPJ(context.Background(), "123")
		if err != nil {
			t.Errorf("Não esperava erro, recebeu: %v", err)
		}
		if client == nil || client.CNPJ != "123" {
			t.Errorf("Esperava cliente com CNPJ 123, recebeu: %v", client)
		}
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Deve retornar cliente pelo ID", func(t *testing.T) {
		clientID := uuid.New()
		expectedClient := &models.Cliente{ID: clientID, CNPJ: "123", NomeFantasia: "Cliente 1"}
		mockRepo := &mockClientRepository{
			mockGetByID: func(id uuid.UUID) (*models.Cliente, error) {
				return expectedClient, nil
			},
		}
		service := NewClientService(mockRepo)

		client, err := service.GetByID(context.Background(), clientID)
		if err != nil {
			t.Errorf("Não esperava erro, recebeu: %v", err)
		}
		if client == nil || client.ID != clientID {
			t.Errorf("Esperava cliente com ID %v, recebeu: %v", clientID, client)
		}
	})
}

func TestUpdateCliente(t *testing.T) {
	t.Run("Deve atualizar cliente com sucesso", func(t *testing.T) {
		mockRepo := &mockClientRepository{
			mockUpdate: func(cliente *models.Cliente) error {
				return nil
			},
		}
		service := NewClientService(mockRepo)

		err := service.UpdateCliente(context.Background(), &models.Cliente{CNPJ: "123", NomeFantasia: "Novo Nome"})
		if err != nil {
			t.Errorf("Não esperava erro, recebeu: %v", err)
		}
	})
}

func TestDeleteCliente(t *testing.T) {
	t.Run("Deve deletar cliente com sucesso", func(t *testing.T) {
		clientID := uuid.New()
		mockRepo := &mockClientRepository{
			mockDelete: func(id uuid.UUID) error {
				if id != clientID {
					t.Errorf("Esperava ID %v, recebeu %v", clientID, id)
				}
				return nil
			},
		}
		service := NewClientService(mockRepo)

		err := service.DeleteCliente(context.Background(), clientID)
		if err != nil {
			t.Errorf("Não esperava erro, recebeu: %v", err)
		}
	})
}
