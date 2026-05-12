package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

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

type AnalistaPostgres struct {
	db *sql.DB
}

func NewAnalistaPostgres(db *sql.DB) *AnalistaPostgres {
	return &AnalistaPostgres{db: db}
}

func (r *AnalistaPostgres) Create(ctx context.Context, a *models.Analista) error {
	query := `INSERT INTO analistas (id, nome, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	now := time.Now()
	_, err := r.db.ExecContext(ctx, query, a.ID, a.Nome, a.Email, now, now)
	return err
}

func (r *AnalistaPostgres) GetByEmail(ctx context.Context, email string) (*models.Analista, error) {
	query := `SELECT id, nome, email, created_at, updated_at, deleted_at FROM analistas WHERE email = $1 AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, email)

	var a models.Analista
	err := row.Scan(&a.ID, &a.Nome, &a.Email, &a.CreatedAt, &a.UpdatedAt, &a.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *AnalistaPostgres) GetByID(ctx context.Context, id uuid.UUID) (*models.Analista, error) {
	query := `SELECT id, nome, email, created_at, updated_at, deleted_at FROM analistas WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, id)

	var a models.Analista
	err := row.Scan(&a.ID, &a.Nome, &a.Email, &a.CreatedAt, &a.UpdatedAt, &a.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *AnalistaPostgres) GetEveryone(ctx context.Context) ([]*models.Analista, error) {
	query := `SELECT id, nome, email, created_at, updated_at, deleted_at FROM analistas WHERE deleted_at IS NULL ORDER BY nome ASC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var analistas []*models.Analista
	for rows.Next() {
		var a models.Analista
		if err := rows.Scan(&a.ID, &a.Nome, &a.Email, &a.CreatedAt, &a.UpdatedAt, &a.DeletedAt); err != nil {
			return nil, err
		}
		analistas = append(analistas, &a)
	}
	return analistas, nil
}

func (r *AnalistaPostgres) Update(ctx context.Context, a *models.Analista) error {
	query := `UPDATE analistas SET nome = $1, email = $2, updated_at = $3 WHERE id = $4 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query, a.Nome, a.Email, time.Now(), a.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("analista não encontrado")
	}
	return nil
}

func (r *AnalistaPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE analistas SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query, time.Now(), id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("analista não encontrado")
	}
	return nil
}
