package service

import (
	"context"
	"errors"

	"github.com/leonardosth/onboardly/internal/models"
	"github.com/leonardosth/onboardly/internal/repository"

	"github.com/google/uuid"
)

type UsuarioService struct {
	repo repository.UsuarioRepository
}

func NewUsuarioService(repo repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

func (s *UsuarioService) Create(ctx context.Context, u *models.Usuario) error {
	existente, err := s.repo.GetByEmail(ctx, u.Email)
	if err != nil {
		return err
	}
	if existente != nil {
		return errors.New("usuário com esse email já cadastrado")
	}
	u.ID = uuid.New()
	// Nota: Para criação via Admin, a senha pode ser um padrão ou gerada.
	// Por enquanto, manteremos a lógica simples.
	return s.repo.Create(ctx, u)
}

func (s *UsuarioService) GetByCargo(ctx context.Context, cargo string) ([]*models.Usuario, error) {
	return s.repo.GetByCargo(ctx, cargo)
}

func (s *UsuarioService) GetByID(ctx context.Context, id uuid.UUID) (*models.Usuario, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UsuarioService) Update(ctx context.Context, u *models.Usuario) error {
	return s.repo.Update(ctx, u)
}

func (s *UsuarioService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
