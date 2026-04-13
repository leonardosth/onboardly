package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type ClientPostgres struct {
	db *sql.DB
}

func NewClientPostgres(db *sql.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (r *ClientPostgres) Create(ctx context.Context, c *models.Cliente) error {
	query := `INSERT INTO clientes (id, nome_fantasia, cnpj, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	now := time.Now()
	_, err := r.db.ExecContext(ctx, query, c.ID, c.NomeFantasia, c.CNPJ, now, now)
	return err
}

func (r *ClientPostgres) GetByCNPJ(ctx context.Context, cnpj string) (*models.Cliente, error) {
	query := `SELECT id, nome_fantasia, cnpj, created_at, updated_at FROM clientes WHERE cnpj = $1`
	row := r.db.QueryRowContext(ctx, query, cnpj)

	var c models.Cliente
	err := row.Scan(&c.ID, &c.NomeFantasia, &c.CNPJ, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *ClientPostgres) GetByID(ctx context.Context, id uuid.UUID) (*models.Cliente, error) {
	query := `SELECT id, nome_fantasia, cnpj, created_at, updated_at FROM clientes WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var c models.Cliente
	err := row.Scan(&c.ID, &c.NomeFantasia, &c.CNPJ, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *ClientPostgres) GetEveryone(ctx context.Context) ([]*models.Cliente, error) {
	query := `SELECT id, nome_fantasia, cnpj, created_at, updated_at FROM clientes ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []*models.Cliente
	for rows.Next() {
		var c models.Cliente
		if err := rows.Scan(&c.ID, &c.NomeFantasia, &c.CNPJ, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		clientes = append(clientes, &c)
	}
	return clientes, nil
}

func (r *ClientPostgres) UpdateCliente(ctx context.Context, c *models.Cliente) error {
	query := `UPDATE clientes SET nome_fantasia = $1, updated_at = $2 WHERE id = $3`
	res, err := r.db.ExecContext(ctx, query, c.NomeFantasia, time.Now(), c.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("cliente não encontrado")
	}
	return nil
}

func (r *ClientPostgres) DeleteCliente(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM clientes WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("cliente não encontrado")
	}
	return nil
}
