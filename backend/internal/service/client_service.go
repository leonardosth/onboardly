package service

import (
	"context"
	"errors"
	"onboardly/internal/repository"

	"github.com/google/uuid"
)

type ClientRepository interface {
	Create(ctx context.Context, cliente *repository.Cliente) error
	GetByCNPJ(ctx context.Context, cnpj string) (*repository.Cliente, error)
	GetEveryone(ctx context.Context) ([]*repository.Cliente, error)
	UpdateCliente(ctx context.Context, cliente *repository.Cliente) error
	DeleteCliente(ctx context.Context, id uuid.UUID) error
}

type ClientService struct {
	repo ClientRepository
}

func NewClientService(repo ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) CreateClient(ctx context.Context, cliente *repository.Cliente) error {
	existente, _ := s.repo.GetByCNPJ(ctx, cliente.CNPJ)
	if existente != nil {
		return errors.New("cliente com esse CNPJ já cadastrado")

	}
	return s.repo.Create(ctx, cliente)
}
