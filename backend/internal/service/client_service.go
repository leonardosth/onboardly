package service

import (
	"context"
	"errors"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type ClientRepository interface {
	Create(ctx context.Context, cliente *models.Cliente) error
	GetByCNPJ(ctx context.Context, cnpj string) (*models.Cliente, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Cliente, error)
	GetEveryone(ctx context.Context) ([]*models.Cliente, error)
	UpdateCliente(ctx context.Context, cliente *models.Cliente) error
	DeleteCliente(ctx context.Context, id uuid.UUID) error
}

type ClientService struct {
	repo ClientRepository
}

func NewClientService(repo ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) CreateClient(ctx context.Context, cliente *models.Cliente) error {
	existente, err := s.repo.GetByCNPJ(ctx, cliente.CNPJ)
	if err != nil {
		return err
	}
	if existente != nil {
		return errors.New("cliente com esse CNPJ já cadastrado")
	}
	cliente.ID = uuid.New()
	return s.repo.Create(ctx, cliente)
}

func (s *ClientService) GetEveryone(ctx context.Context) ([]*models.Cliente, error) {
	return s.repo.GetEveryone(ctx)
}

func (s *ClientService) GetByCNPJ(ctx context.Context, cnpj string) (*models.Cliente, error) {
	return s.repo.GetByCNPJ(ctx, cnpj)
}

func (s *ClientService) GetByID(ctx context.Context, id uuid.UUID) (*models.Cliente, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ClientService) UpdateCliente(ctx context.Context, cliente *models.Cliente) error {
	return s.repo.UpdateCliente(ctx, cliente)
}

func (s *ClientService) DeleteCliente(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteCliente(ctx, id)
}
