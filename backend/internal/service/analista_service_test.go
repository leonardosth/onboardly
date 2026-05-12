package service

import (
	"context"
	"testing"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type mockAnalistaRepository struct {
	mockGetByEmail  func(email string) (*models.Analista, error)
	mockGetByID     func(id uuid.UUID) (*models.Analista, error)
	mockCreate      func(a *models.Analista) error
	mockGetEveryone func() ([]*models.Analista, error)
	mockUpdate      func(a *models.Analista) error
	mockDelete      func(id uuid.UUID) error
}

func (m *mockAnalistaRepository) GetByEmail(ctx context.Context, email string) (*models.Analista, error) {
	return m.mockGetByEmail(email)
}
func (m *mockAnalistaRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Analista, error) {
	return m.mockGetByID(id)
}
func (m *mockAnalistaRepository) Create(ctx context.Context, a *models.Analista) error {
	return m.mockCreate(a)
}
func (m *mockAnalistaRepository) GetEveryone(ctx context.Context) ([]*models.Analista, error) {
	return m.mockGetEveryone()
}
func (m *mockAnalistaRepository) Update(ctx context.Context, a *models.Analista) error {
	return m.mockUpdate(a)
}
func (m *mockAnalistaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return m.mockDelete(id)
}

func TestCreateAnalista(t *testing.T) {
	t.Run("Deve retornar erro se Email ja existir", func(t *testing.T) {
		mockRepo := &mockAnalistaRepository{
			mockGetByEmail: func(email string) (*models.Analista, error) {
				return &models.Analista{Email: "test@test.com"}, nil
			},
		}
		service := NewAnalistaService(mockRepo)

		err := service.Create(context.Background(), &models.Analista{Email: "test@test.com"})
		if err == nil || err.Error() != "analista com esse email já cadastrado" {
			t.Errorf("Esperava erro de email duplicado, recebeu: %v", err)
		}
	})
}
