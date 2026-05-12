package service

import (
	"context"
	"errors"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type AnalistaRepository interface {
	Create(ctx context.Context, a *models.Analista) error
	GetByEmail(ctx context.Context, email string) (*models.Analista, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Analista, error)
	GetEveryone(ctx context.Context) ([]*models.Analista, error)
	Update(ctx context.Context, a *models.Analista) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type AnalistaService struct {
	repo AnalistaRepository
}

func NewAnalistaService(repo AnalistaRepository) *AnalistaService {
	return &AnalistaService{repo: repo}
}

func (s *AnalistaService) Create(ctx context.Context, a *models.Analista) error {
	existente, err := s.repo.GetByEmail(ctx, a.Email)
	if err != nil {
		return err
	}
	if existente != nil {
		return errors.New("analista com esse email já cadastrado")
	}
	a.ID = uuid.New()
	return s.repo.Create(ctx, a)
}

func (s *AnalistaService) GetEveryone(ctx context.Context) ([]*models.Analista, error) {
	return s.repo.GetEveryone(ctx)
}

func (s *AnalistaService) GetByID(ctx context.Context, id uuid.UUID) (*models.Analista, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *AnalistaService) Update(ctx context.Context, a *models.Analista) error {
	return s.repo.Update(ctx, a)
}

func (s *AnalistaService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
