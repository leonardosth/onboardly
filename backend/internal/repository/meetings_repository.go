package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Meeting struct {
	ID            uuid.UUID `json:"id"`
	ProjetoID     uuid.UUID `json:"projeto_id"`
	DataAgendada  string    `json:"data_agendada"`
	StatusReuniao string    `json:"status_reuniao"`
	Observacoes   string    `json:"observacoes"`
}

type MeetingRepository struct {
	db *sql.DB
}

func NewMeetingRepository(db *sql.DB) *MeetingRepository {
	return &MeetingRepository{db: db}
}

func (r *MeetingRepository) Create(ctx context.Context, meeting *Meeting) error {
	query := `
		INSERT INTO meetings (id, projeto_id, data_agendada, status_reuniao, observacoes)
		VALUES ($1, $2, $3, $4, $5)`

	meeting.ID = uuid.New()

	_, err := r.db.ExecContext(ctx, query,
		meeting.ID,
		meeting.ProjetoID,
		meeting.DataAgendada,
		meeting.StatusReuniao,
		meeting.Observacoes,
	)
	if err != nil {
		return fmt.Errorf("failed to insert meeting: %w", err)
	}
	return nil
}

func (r *MeetingRepository) GetByID(ctx context.Context, id uuid.UUID) (*Meeting, error) {
	query := `SELECT id, projeto_id, data_agendada, status_reuniao, observacoes FROM meetings WHERE id = $1`

	meeting := &Meeting{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&meeting.ID,
		&meeting.ProjetoID,
		&meeting.DataAgendada,
		&meeting.StatusReuniao,
		&meeting.Observacoes,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query meeting by id: %w", err)
	}
	return meeting, nil
}

func (r *MeetingRepository) GetByCliente(ctx context.Context, clienteID uuid.UUID) ([]*Meeting, error) {
	query := `
		SELECT m.id, m.projeto_id, m.data_agendada, m.status_reuniao, m.observacoes
		FROM meetings m
		JOIN projetos p ON m.projeto_id = p.id
		WHERE p.cliente_id = $1`

	rows, err := r.db.QueryContext(ctx, query, clienteID)
	if err != nil {
		return nil, fmt.Errorf("failed to query meetings by cliente: %w", err)
	}
	defer rows.Close()

	var meetings []*Meeting
	for rows.Next() {
		meeting := &Meeting{}
		err := rows.Scan(
			&meeting.ID,
			&meeting.ProjetoID,
			&meeting.DataAgendada,
			&meeting.StatusReuniao,
			&meeting.Observacoes,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan meeting: %w", err)
		}
		meetings = append(meetings, meeting)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return meetings, nil
}

func (r *MeetingRepository) GetByAnalista(ctx context.Context, analistaID uuid.UUID) ([]*Meeting, error) {
	query := `
		SELECT m.id, m.projeto_id, m.data_agendada, m.status_reuniao, m.observacoes
		FROM meetings m
		JOIN projetos p ON m.projeto_id = p.id
		WHERE p.analista_id = $1`

	rows, err := r.db.QueryContext(ctx, query, analistaID)
	if err != nil {
		return nil, fmt.Errorf("failed to query meetings by analista: %w", err)
	}
	defer rows.Close()

	var meetings []*Meeting
	for rows.Next() {
		meeting := &Meeting{}
		err := rows.Scan(
			&meeting.ID,
			&meeting.ProjetoID,
			&meeting.DataAgendada,
			&meeting.StatusReuniao,
			&meeting.Observacoes,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan meeting: %w", err)
		}
		meetings = append(meetings, meeting)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return meetings, nil
}

func (r *MeetingRepository) GetAll(ctx context.Context) ([]*Meeting, error) {
	query := `SELECT id, projeto_id, data_agendada, status_reuniao, observacoes FROM meetings`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all meetings: %w", err)
	}
	defer rows.Close()

	var meetings []*Meeting
	for rows.Next() {
		meeting := &Meeting{}
		err := rows.Scan(
			&meeting.ID,
			&meeting.ProjetoID,
			&meeting.DataAgendada,
			&meeting.StatusReuniao,
			&meeting.Observacoes,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan meeting: %w", err)
		}
		meetings = append(meetings, meeting)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return meetings, nil
}

func (r *MeetingRepository) Update(ctx context.Context, meeting *Meeting) error {
	query := `
		UPDATE meetings
		SET projeto_id = $1, data_agendada = $2, status_reuniao = $3, observacoes = $4
		WHERE id = $5`

	_, err := r.db.ExecContext(ctx, query,
		meeting.ProjetoID,
		meeting.DataAgendada,
		meeting.StatusReuniao,
		meeting.Observacoes,
		meeting.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update meeting: %w", err)
	}
	return nil
}

func (r *MeetingRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM meetings WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete meeting: %w", err)
	}
	return nil
}
