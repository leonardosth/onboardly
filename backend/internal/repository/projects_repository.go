package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Projeto struct {
	ID              uuid.UUID `json:"id"`
	ClienteID       uuid.UUID `json:"cliente_id"`
	AnalistaID      uuid.UUID `json:"analista_id"`
	DataContratacao string    `json:"data_contratacao"`
	DataAtivacao    string    `json:"data_ativacao"`
	StatusAtivacao  string    `json:"status_ativacao"`
	StatusProjeto   string    `json:"status_projeto"`
}

type ProjetoRepository struct {
	db *sql.DB
}

func NewProjetoRepository(db *sql.DB) *ProjetoRepository {
	return &ProjetoRepository{db: db}
}

func (r *ProjetoRepository) Create(projeto *Projeto) error {
	query := `
		INSERT INTO projetos (id, cliente_id, analista_id, data_contratacao, data_ativacao, status_ativacao, status_projeto)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	projeto.ID = uuid.New()

	_, err := r.db.Exec(query,
		projeto.ID,
		projeto.ClienteID,
		projeto.AnalistaID,
		projeto.DataContratacao,
		projeto.DataAtivacao,
		projeto.StatusAtivacao,
		projeto.StatusProjeto)

	if err != nil {
		return err
	}
	return nil
}

func (r *ProjetoRepository) GetByID(ctx context.Context, id uuid.UUID) (*Projeto, error) {
	query := `SELECT id, cliente_id, analista_id, data_contratacao, data_ativacao, status_ativacao, status_projeto FROM projetos WHERE id = $1`

	projeto := &Projeto{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&projeto.ID,
		&projeto.ClienteID,
		&projeto.AnalistaID,
		&projeto.DataContratacao,
		&projeto.DataAtivacao,
		&projeto.StatusAtivacao,
		&projeto.StatusProjeto,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query projeto by id: %w", err)
	}

	return projeto, nil
}

func (r *ProjetoRepository) GetByCliente(ctx context.Context, clienteID uuid.UUID) ([]*Projeto, error) {
	query := `SELECT id, cliente_id, analista_id, data_contratacao, data_ativacao, status_ativacao, status_projeto FROM projetos WHERE cliente_id = $1`

	rows, err := r.db.QueryContext(ctx, query, clienteID)
	if err != nil {
		return nil, fmt.Errorf("failed to query projetos by cliente: %w", err)
	}
	defer rows.Close()

	var projetos []*Projeto
	for rows.Next() {
		projeto := &Projeto{}
		err := rows.Scan(
			&projeto.ID,
			&projeto.ClienteID,
			&projeto.AnalistaID,
			&projeto.DataContratacao,
			&projeto.DataAtivacao,
			&projeto.StatusAtivacao,
			&projeto.StatusProjeto,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan projeto: %w", err)
		}
		projetos = append(projetos, projeto)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return projetos, nil
}

func (r *ProjetoRepository) GetByAnalista(ctx context.Context, analistaID uuid.UUID) ([]*Projeto, error) {
	query := `SELECT id, cliente_id, analista_id, data_contratacao, data_ativacao, status_ativacao, status_projeto FROM projetos WHERE analista_id = $1`

	rows, err := r.db.QueryContext(ctx, query, analistaID)
	if err != nil {
		return nil, fmt.Errorf("failed to query projetos by analista: %w", err)
	}
	defer rows.Close()

	var projetos []*Projeto
	for rows.Next() {
		projeto := &Projeto{}
		err := rows.Scan(
			&projeto.ID,
			&projeto.ClienteID,
			&projeto.AnalistaID,
			&projeto.DataContratacao,
			&projeto.DataAtivacao,
			&projeto.StatusAtivacao,
			&projeto.StatusProjeto,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan projeto: %w", err)
		}
		projetos = append(projetos, projeto)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return projetos, nil
}

func (r *ProjetoRepository) GetAll(ctx context.Context) ([]*Projeto, error) {
	query := `SELECT id, cliente_id, analista_id, data_contratacao, data_ativacao, status_ativacao, status_projeto FROM projetos`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all projetos: %w", err)
	}
	defer rows.Close()

	var projetos []*Projeto
	for rows.Next() {
		projeto := &Projeto{}
		err := rows.Scan(
			&projeto.ID,
			&projeto.ClienteID,
			&projeto.AnalistaID,
			&projeto.DataContratacao,
			&projeto.DataAtivacao,
			&projeto.StatusAtivacao,
			&projeto.StatusProjeto,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan projeto: %w", err)
		}
		projetos = append(projetos, projeto)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return projetos, nil
}

func (r *ProjetoRepository) Update(ctx context.Context, projeto *Projeto) error {
	query := `
		UPDATE projetos
		SET cliente_id = $1, analista_id = $2, data_contratacao = $3, data_ativacao = $4, status_ativacao = $5, status_projeto = $6
		WHERE id = $7`

	_, err := r.db.ExecContext(ctx, query,
		projeto.ClienteID,
		projeto.AnalistaID,
		projeto.DataContratacao,
		projeto.DataAtivacao,
		projeto.StatusAtivacao,
		projeto.StatusProjeto,
		projeto.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update projeto: %w", err)
	}
	return nil
}

func (r *ProjetoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM projetos WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete projeto: %w", err)
	}
	return nil
}
