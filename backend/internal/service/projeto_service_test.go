package service

import (
	"context"
	"testing"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type mockProjetoRepository struct {
	mockGetByID     func(id uuid.UUID) (*models.Projeto, error)
	mockGetByCliente func(id uuid.UUID) ([]*models.Projeto, error)
	mockCreate      func(p *models.Projeto) error
	mockGetEveryone func() ([]*models.Projeto, error)
	mockUpdate      func(p *models.Projeto) error
	mockDelete      func(id uuid.UUID) error
	mockGetDashboardStats func() (*models.DashboardStats, error)
}

func (m *mockProjetoRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Projeto, error) {
	return m.mockGetByID(id)
}
func (m *mockProjetoRepository) GetByCliente(ctx context.Context, id uuid.UUID) ([]*models.Projeto, error) {
	return m.mockGetByCliente(id)
}
func (m *mockProjetoRepository) Create(ctx context.Context, p *models.Projeto) error {
	return m.mockCreate(p)
}
func (m *mockProjetoRepository) GetEveryone(ctx context.Context) ([]*models.Projeto, error) {
	return m.mockGetEveryone()
}
func (m *mockProjetoRepository) Update(ctx context.Context, p *models.Projeto) error {
	return m.mockUpdate(p)
}
func (m *mockProjetoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return m.mockDelete(id)
}
func (m *mockProjetoRepository) GetDashboardStats(ctx context.Context) (*models.DashboardStats, error) {
	return m.mockGetDashboardStats()
}

func TestCreateProjeto(t *testing.T) {
	t.Run("Deve criar projeto com status inicial Backlog", func(t *testing.T) {
		mockRepo := &mockProjetoRepository{
			mockCreate: func(p *models.Projeto) error {
				if p.StatusProjeto != models.StatusBacklog {
					t.Errorf("Esperava status Backlog, recebeu %v", p.StatusProjeto)
				}
				return nil
			},
		}
		service := NewProjetoService(mockRepo)
		err := service.Create(context.Background(), &models.Projeto{})
		if err != nil {
			t.Errorf("Nao esperava erro, recebeu: %v", err)
		}
	})
}
