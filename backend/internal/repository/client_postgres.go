package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type ClientRepository interface {
	Create(ctx context.Context, c *models.Cliente) error
	GetByCNPJ(ctx context.Context, cnpj string) (*models.Cliente, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Cliente, error)
	GetEveryone(ctx context.Context) ([]*models.Cliente, error)
	Update(ctx context.Context, c *models.Cliente) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ClientPostgres struct {
	db *sql.DB
}

func NewClientPostgres(db *sql.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (r *ClientPostgres) Create(ctx context.Context, c *models.Cliente) error {
	query := `INSERT INTO clientes (id, nome, cnpj, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	now := time.Now()
	_, err := r.db.ExecContext(ctx, query, c.ID, c.Nome, c.CNPJ, now, now)
	return err
}

func (r *ClientPostgres) GetByCNPJ(ctx context.Context, cnpj string) (*models.Cliente, error) {
	query := `SELECT id, nome, cnpj, created_at, updated_at, deleted_at FROM clientes WHERE cnpj = $1 AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, cnpj)

	var c models.Cliente
	err := row.Scan(&c.ID, &c.Nome, &c.CNPJ, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *ClientPostgres) GetByID(ctx context.Context, id uuid.UUID) (*models.Cliente, error) {
	query := `SELECT id, nome, cnpj, created_at, updated_at, deleted_at FROM clientes WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, id)

	var c models.Cliente
	err := row.Scan(&c.ID, &c.Nome, &c.CNPJ, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *ClientPostgres) GetEveryone(ctx context.Context) ([]*models.Cliente, error) {
	query := `SELECT id, nome, cnpj, created_at, updated_at, deleted_at FROM clientes WHERE deleted_at IS NULL ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []*models.Cliente
	for rows.Next() {
		var c models.Cliente
		if err := rows.Scan(&c.ID, &c.Nome, &c.CNPJ, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt); err != nil {
			return nil, err
		}
		clientes = append(clientes, &c)
	}
	return clientes, nil
}

func (r *ClientPostgres) Update(ctx context.Context, c *models.Cliente) error {
	query := `UPDATE clientes SET nome = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query, c.Nome, time.Now(), c.ID)
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

func (r *ClientPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE clientes SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query, time.Now(), id)
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
