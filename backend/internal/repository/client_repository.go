package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Cliente struct {
	ID           uuid.UUID `json:"id"`
	NomeFantasia string    `json:"nome_fantasia"`
	CNPJ         string    `json:"cnpj"`
}

type ClienteRepository struct {
	db *sql.DB
}

func NewClienteRepository(db *sql.DB) *ClienteRepository {
	return &ClienteRepository{db: db}
}

func (r *ClienteRepository) Create(ctx context.Context, cliente *Cliente) error {
	query := `
		INSERT INTO clientes (id, nome_fantasia, cnpj)
		VALUES ($1, $2, $3)`

	cliente.ID = uuid.New()

	_, err := r.db.ExecContext(ctx, query,
		cliente.ID,
		cliente.NomeFantasia,
		cliente.CNPJ,
	)

	if err != nil {
		return fmt.Errorf("failed to insert cliente: %w", err)
	}

	return nil
}

func (r *ClienteRepository) GetByCNPJ(ctx context.Context, cnpj string) (*Cliente, error) {
	query := `SELECT id, nome_fantasia, cnpj FROM clientes WHERE cnpj = $1`

	cliente := &Cliente{}
	err := r.db.QueryRowContext(ctx, query, cnpj).Scan(
		&cliente.ID,
		&cliente.NomeFantasia,
		&cliente.CNPJ,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query cliente: %w", err)
	}

	return cliente, nil
}

func (r *ClienteRepository) GetEveryone(ctx context.Context) ([]*Cliente, error) {
	query := `SELECT id, nome_fantasia, cnpj FROM clientes`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query clientes: %w", err)
	}
	defer rows.Close()

	var clientes []*Cliente
	for rows.Next() {
		cliente := &Cliente{}
		err := rows.Scan(
			&cliente.ID,
			&cliente.NomeFantasia,
			&cliente.CNPJ,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cliente: %w", err)
		}
		clientes = append(clientes, cliente)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return clientes, nil
}

func (r *ClienteRepository) UpdateCliente(ctx context.Context, cliente *Cliente) error {
	query := `
		UPDATE clientes
		SET nome_fantasia = $1, cnpj = $2
		WHERE id = $3`

	_, err := r.db.ExecContext(ctx, query, cliente.NomeFantasia, cliente.CNPJ, cliente.ID)
	return err
}

func (r *ClienteRepository) DeleteCliente(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM clientes WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
