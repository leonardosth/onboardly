package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Analista struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Email string    `json:"email"`
}

type AnalistaRepository struct {
	db *sql.DB
}

func NewAnalistaRepository(db *sql.DB) *AnalistaRepository {
	return &AnalistaRepository{db: db}
}

func (r *AnalistaRepository) Create(ctx context.Context, analista *Analista) error {
	query := `
		INSERT INTO analistas (id, nome, email)
		VALUES ($1, $2, $3)`

	analista.ID = uuid.New()

	_, err := r.db.ExecContext(ctx, query,
		analista.ID,
		analista.Nome,
		analista.Email,
	)

	if err != nil {
		return fmt.Errorf("failed to insert analista: %w", err)
	}

	return nil
}

func (r *AnalistaRepository) GetByEmail(ctx context.Context, email string) (*Analista, error) {
	query := `SELECT id, nome, email FROM analistas WHERE email = $1`

	analista := &Analista{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&analista.ID,
		&analista.Nome,
		&analista.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query analista: %w", err)
	}

	return analista, nil
}

func (r *AnalistaRepository) GetEveryone(ctx context.Context) ([]*Analista, error) {
	query := `SELECT id, nome, email FROM analistas`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query analistas: %w", err)
	}
	defer rows.Close()

	var analistas []*Analista
	for rows.Next() {
		analista := &Analista{}
		err := rows.Scan(
			&analista.ID,
			&analista.Nome,
			&analista.Email,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan analista: %w", err)
		}
		analistas = append(analistas, analista)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return analistas, nil
}

func (r *AnalistaRepository) UpdateAnalista(ctx context.Context, analista *Analista) error {
	query := `
		UPDATE analistas
		SET nome = $1, email = $2
		WHERE id = $3`

	_, err := r.db.ExecContext(ctx, query, analista.Nome, analista.Email, analista.ID)
	return err
}

func (r *AnalistaRepository) DeleteAnalista(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM analistas WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
