package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

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

type ReuniaoPostgres struct {
	db *sql.DB
}

func NewReuniaoPostgres(db *sql.DB) *ReuniaoPostgres {
	return &ReuniaoPostgres{db: db}
}

func (r *ReuniaoPostgres) Create(ctx context.Context, reuniao *models.Reuniao) error {
	query := `INSERT INTO reunioes (id, projeto_id, analista_id, data_agendada, status, observacoes, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	now := time.Now()
	_, err := r.db.ExecContext(ctx, query, reuniao.ID, reuniao.ProjetoID, reuniao.AnalistaID, reuniao.DataAgendada, reuniao.Status, reuniao.Observacoes, now, now)
	return err
}

func (r *ReuniaoPostgres) GetByID(ctx context.Context, id uuid.UUID) (*models.Reuniao, error) {
	query := `SELECT id, projeto_id, analista_id, data_agendada, status, observacoes, created_at, updated_at, deleted_at 
	          FROM reunioes WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, id)

	var reuniao models.Reuniao
	err := row.Scan(&reuniao.ID, &reuniao.ProjetoID, &reuniao.AnalistaID, &reuniao.DataAgendada, &reuniao.Status, &reuniao.Observacoes, &reuniao.CreatedAt, &reuniao.UpdatedAt, &reuniao.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &reuniao, nil
}

func (r *ReuniaoPostgres) GetByProjeto(ctx context.Context, projetoID uuid.UUID) ([]*models.Reuniao, error) {
	query := `SELECT id, projeto_id, analista_id, data_agendada, status, observacoes, created_at, updated_at, deleted_at 
	          FROM reunioes WHERE projeto_id = $1 AND deleted_at IS NULL`
	rows, err := r.db.QueryContext(ctx, query, projetoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reunioes []*models.Reuniao
	for rows.Next() {
		var reuniao models.Reuniao
		if err := rows.Scan(&reuniao.ID, &reuniao.ProjetoID, &reuniao.AnalistaID, &reuniao.DataAgendada, &reuniao.Status, &reuniao.Observacoes, &reuniao.CreatedAt, &reuniao.UpdatedAt, &reuniao.DeletedAt); err != nil {
			return nil, err
		}
		reunioes = append(reunioes, &reuniao)
	}
	return reunioes, nil
}

func (r *ReuniaoPostgres) GetEveryone(ctx context.Context) ([]*models.Reuniao, error) {
	query := `SELECT id, projeto_id, analista_id, data_agendada, status, observacoes, created_at, updated_at, deleted_at 
	          FROM reunioes WHERE deleted_at IS NULL ORDER BY data_agendada DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reunioes []*models.Reuniao
	for rows.Next() {
		var reuniao models.Reuniao
		if err := rows.Scan(&reuniao.ID, &reuniao.ProjetoID, &reuniao.AnalistaID, &reuniao.DataAgendada, &reuniao.Status, &reuniao.Observacoes, &reuniao.CreatedAt, &reuniao.UpdatedAt, &reuniao.DeletedAt); err != nil {
			return nil, err
		}
		reunioes = append(reunioes, &reuniao)
	}
	return reunioes, nil
}

func (r *ReuniaoPostgres) Update(ctx context.Context, reuniao *models.Reuniao) error {
	query := `UPDATE reunioes SET projeto_id = $1, analista_id = $2, data_agendada = $3, status = $4, observacoes = $5, updated_at = $6 
	          WHERE id = $7 AND deleted_at IS NULL`
	_, err := r.db.ExecContext(ctx, query, reuniao.ProjetoID, reuniao.AnalistaID, reuniao.DataAgendada, reuniao.Status, reuniao.Observacoes, time.Now(), reuniao.ID)
	return err
}

func (r *ReuniaoPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE reunioes SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query, time.Now(), id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("reunião não encontrada")
	}
	return nil
}
