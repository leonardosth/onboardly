package service

import (
	"context"
	"testing"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type mockUsuarioRepository struct {
	mockGetByEmail func(email string) (*models.Usuario, error)
	mockGetByID    func(id uuid.UUID) (*models.Usuario, error)
	mockCreate     func(u *models.Usuario) error
	mockGetByCargo func(cargo string) ([]*models.Usuario, error)
	mockUpdate     func(u *models.Usuario) error
	mockDelete     func(id uuid.UUID) error
}

func (m *mockUsuarioRepository) GetByEmail(ctx context.Context, email string) (*models.Usuario, error) {
	return m.mockGetByEmail(email)
}
func (m *mockUsuarioRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Usuario, error) {
	return m.mockGetByID(id)
}
func (m *mockUsuarioRepository) Create(ctx context.Context, u *models.Usuario) error {
	return m.mockCreate(u)
}
func (m *mockUsuarioRepository) GetByCargo(ctx context.Context, cargo string) ([]*models.Usuario, error) {
	return m.mockGetByCargo(cargo)
}
func (m *mockUsuarioRepository) Update(ctx context.Context, u *models.Usuario) error {
	return m.mockUpdate(u)
}
func (m *mockUsuarioRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return m.mockDelete(id)
}

func TestCreateUsuario(t *testing.T) {
	t.Run("Deve retornar erro se Email ja existir", func(t *testing.T) {
		mockRepo := &mockUsuarioRepository{
			mockGetByEmail: func(email string) (*models.Usuario, error) {
				return &models.Usuario{Email: "test@test.com"}, nil
			},
		}
		service := NewUsuarioService(mockRepo)

		err := service.Create(context.Background(), &models.Usuario{Email: "test@test.com"})
		if err == nil || err.Error() != "usuário com esse email já cadastrado" {
			t.Errorf("Esperava erro de email duplicado, recebeu: %v", err)
		}
	})
}
