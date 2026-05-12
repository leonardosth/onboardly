package service

import (
	"context"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type ReuniaoRepository interface {
	Create(ctx context.Context, r *models.Reuniao) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Reuniao, error)
	GetByProjeto(ctx context.Context, projetoID uuid.UUID) ([]*models.Reuniao, error)
	GetEveryone(ctx context.Context) ([]*models.Reuniao, error)
	Update(ctx context.Context, r *models.Reuniao) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ReuniaoService struct {
	repo ReuniaoRepository
}

func NewReuniaoService(repo ReuniaoRepository) *ReuniaoService {
	return &ReuniaoService{repo: repo}
}

func (s *ReuniaoService) Create(ctx context.Context, r *models.Reuniao) error {
	r.ID = uuid.New()
	if r.Status == "" {
		r.Status = models.StatusAgendada
	}
	return s.repo.Create(ctx, r)
}

func (s *ReuniaoService) GetByID(ctx context.Context, id uuid.UUID) (*models.Reuniao, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ReuniaoService) GetEveryone(ctx context.Context) ([]*models.Reuniao, error) {
	return s.repo.GetEveryone(ctx)
}

func (s *ReuniaoService) Update(ctx context.Context, r *models.Reuniao) error {
	return s.repo.Update(ctx, r)
}

func (s *ReuniaoService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
