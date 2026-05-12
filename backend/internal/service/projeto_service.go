package service

import (
	"context"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type ProjetoRepository interface {
	Create(ctx context.Context, p *models.Projeto) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Projeto, error)
	GetByCliente(ctx context.Context, clienteID uuid.UUID) ([]*models.Projeto, error)
	GetEveryone(ctx context.Context) ([]*models.Projeto, error)
	Update(ctx context.Context, p *models.Projeto) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetDashboardStats(ctx context.Context) (*models.DashboardStats, error)
}

type ProjetoService struct {
	repo ProjetoRepository
}

func NewProjetoService(repo ProjetoRepository) *ProjetoService {
	return &ProjetoService{repo: repo}
}

func (s *ProjetoService) Create(ctx context.Context, p *models.Projeto) error {
	p.ID = uuid.New()
	if p.StatusProjeto == "" {
		p.StatusProjeto = models.StatusBacklog
	}
	return s.repo.Create(ctx, p)
}

func (s *ProjetoService) GetByID(ctx context.Context, id uuid.UUID) (*models.Projeto, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProjetoService) GetEveryone(ctx context.Context) ([]*models.Projeto, error) {
	return s.repo.GetEveryone(ctx)
}

func (s *ProjetoService) Update(ctx context.Context, p *models.Projeto) error {
	return s.repo.Update(ctx, p)
}

func (s *ProjetoService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProjetoService) GetStats(ctx context.Context) (*models.DashboardStats, error) {
	return s.repo.GetDashboardStats(ctx)
}
