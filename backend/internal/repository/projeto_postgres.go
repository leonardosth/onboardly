package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/leonardosth/onboardly/internal/models"

	"github.com/google/uuid"
)

type ProjetoRepository interface {
	Create(ctx context.Context, p *models.Projeto) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Projeto, error)
	GetByCliente(ctx context.Context, clienteID uuid.UUID) ([]*models.Projeto, error)
	GetEveryone(ctx context.Context) ([]*models.Projeto, error)
	Update(ctx context.Context, p *models.Projeto) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetDashboardStats(ctx context.Context) (*models.DashboardStats, error)
}

type ProjetoPostgres struct {
	db *sql.DB
}

func NewProjetoPostgres(db *sql.DB) *ProjetoPostgres {
	return &ProjetoPostgres{db: db}
}

func (r *ProjetoPostgres) Create(ctx context.Context, p *models.Projeto) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO projetos_implantacao (id, cliente_id, analista_id, data_contratacao, data_ativacao, status_ativacao, status_projeto, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	now := time.Now()
	_, err = tx.ExecContext(ctx, query, p.ID, p.ClienteID, p.AnalistaID, p.DataContratacao, p.DataAtivacao, p.StatusAtivacao, p.StatusProjeto, now, now)
	if err != nil {
		return err
	}

	// Histórico inicial
	histQuery := `INSERT INTO projeto_status_historico (id, projeto_id, status_novo, created_at) VALUES ($1, $2, $3, $4)`
	_, err = tx.ExecContext(ctx, histQuery, uuid.New(), p.ID, p.StatusProjeto, now)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ProjetoPostgres) GetByID(ctx context.Context, id uuid.UUID) (*models.Projeto, error) {
	query := `SELECT id, cliente_id, analista_id, data_contratacao, data_ativacao, status_ativacao, status_projeto, created_at, updated_at, deleted_at 
	          FROM projetos_implantacao WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, id)

	var p models.Projeto
	err := row.Scan(&p.ID, &p.ClienteID, &p.AnalistaID, &p.DataContratacao, &p.DataAtivacao, &p.StatusAtivacao, &p.StatusProjeto, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *ProjetoPostgres) GetByCliente(ctx context.Context, clienteID uuid.UUID) ([]*models.Projeto, error) {
	query := `SELECT id, cliente_id, analista_id, data_contratacao, data_ativacao, status_ativacao, status_projeto, created_at, updated_at, deleted_at 
	          FROM projetos_implantacao WHERE cliente_id = $1 AND deleted_at IS NULL`
	rows, err := r.db.QueryContext(ctx, query, clienteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projetos []*models.Projeto
	for rows.Next() {
		var p models.Projeto
		if err := rows.Scan(&p.ID, &p.ClienteID, &p.AnalistaID, &p.DataContratacao, &p.DataAtivacao, &p.StatusAtivacao, &p.StatusProjeto, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
			return nil, err
		}
		projetos = append(projetos, &p)
	}
	return projetos, nil
}

func (r *ProjetoPostgres) GetEveryone(ctx context.Context) ([]*models.Projeto, error) {
	query := `SELECT id, cliente_id, analista_id, data_contratacao, data_ativacao, status_ativacao, status_projeto, created_at, updated_at, deleted_at 
	          FROM projetos_implantacao WHERE deleted_at IS NULL ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projetos []*models.Projeto
	for rows.Next() {
		var p models.Projeto
		if err := rows.Scan(&p.ID, &p.ClienteID, &p.AnalistaID, &p.DataContratacao, &p.DataAtivacao, &p.StatusAtivacao, &p.StatusProjeto, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
			return nil, err
		}
		projetos = append(projetos, &p)
	}
	return projetos, nil
}

func (r *ProjetoPostgres) Update(ctx context.Context, p *models.Projeto) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Buscar status atual para ver se mudou
	var currentStatus models.StatusProjeto
	err = tx.QueryRowContext(ctx, "SELECT status_projeto FROM projetos_implantacao WHERE id = $1 AND deleted_at IS NULL", p.ID).Scan(&currentStatus)
	if err != nil {
		return err
	}

	query := `UPDATE projetos_implantacao SET cliente_id = $1, analista_id = $2, data_contratacao = $3, data_ativacao = $4, status_ativacao = $5, status_projeto = $6, updated_at = $7 
	          WHERE id = $8 AND deleted_at IS NULL`
	now := time.Now()
	_, err = tx.ExecContext(ctx, query, p.ClienteID, p.AnalistaID, p.DataContratacao, p.DataAtivacao, p.StatusAtivacao, p.StatusProjeto, now, p.ID)
	if err != nil {
		return err
	}

	if currentStatus != p.StatusProjeto {
		histQuery := `INSERT INTO projeto_status_historico (id, projeto_id, status_antigo, status_novo, created_at) VALUES ($1, $2, $3, $4, $5)`
		_, err = tx.ExecContext(ctx, histQuery, uuid.New(), p.ID, currentStatus, p.StatusProjeto, now)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *ProjetoPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE projetos_implantacao SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query, time.Now(), id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("projeto não encontrado")
	}
	return nil
}

func (r *ProjetoPostgres) GetDashboardStats(ctx context.Context) (*models.DashboardStats, error) {
	query := `SELECT status_projeto, COUNT(*) 
	          FROM projetos_implantacao 
	          WHERE deleted_at IS NULL 
	          GROUP BY status_projeto`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := &models.DashboardStats{
		PorStatus: make(map[string]int),
	}

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, err
		}
		stats.PorStatus[status] = count
		stats.TotalProjetos += count
	}

	return stats, nil
}
