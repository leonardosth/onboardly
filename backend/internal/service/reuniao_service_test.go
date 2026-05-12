package service

import (
	"context"
	"testing"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type mockReuniaoRepository struct {
	mockGetByID     func(id uuid.UUID) (*models.Reuniao, error)
	mockGetByProjeto func(id uuid.UUID) ([]*models.Reuniao, error)
	mockCreate      func(r *models.Reuniao) error
	mockGetEveryone func() ([]*models.Reuniao, error)
	mockUpdate      func(r *models.Reuniao) error
	mockDelete      func(id uuid.UUID) error
}

func (m *mockReuniaoRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Reuniao, error) {
	return m.mockGetByID(id)
}
func (m *mockReuniaoRepository) GetByProjeto(ctx context.Context, id uuid.UUID) ([]*models.Reuniao, error) {
	return m.mockGetByProjeto(id)
}
func (m *mockReuniaoRepository) Create(ctx context.Context, r *models.Reuniao) error {
	return m.mockCreate(r)
}
func (m *mockReuniaoRepository) GetEveryone(ctx context.Context) ([]*models.Reuniao, error) {
	return m.mockGetEveryone()
}
func (m *mockReuniaoRepository) Update(ctx context.Context, r *models.Reuniao) error {
	return m.mockUpdate(r)
}
func (m *mockReuniaoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return m.mockDelete(id)
}

func TestCreateReuniao(t *testing.T) {
	t.Run("Deve criar reuniao com status inicial Agendada", func(t *testing.T) {
		mockRepo := &mockReuniaoRepository{
			mockCreate: func(r *models.Reuniao) error {
				if r.Status != models.StatusAgendada {
					t.Errorf("Esperava status Agendada, recebeu %v", r.Status)
				}
				return nil
			},
		}
		service := NewReuniaoService(mockRepo)
		err := service.Create(context.Background(), &models.Reuniao{})
		if err != nil {
			t.Errorf("Nao esperava erro, recebeu: %v", err)
		}
	})
}
