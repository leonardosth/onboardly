package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leonardosth/onboardly/internal/models"
)

type UsuarioRepository interface {
	Create(ctx context.Context, u *models.Usuario) error
	GetByEmail(ctx context.Context, email string) (*models.Usuario, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Usuario, error)
	GetByCargo(ctx context.Context, cargo string) ([]*models.Usuario, error)
	Update(ctx context.Context, u *models.Usuario) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UsuarioPostgres struct {
	db *sql.DB
}

func NewUsuarioPostgres(db *sql.DB) *UsuarioPostgres {
	return &UsuarioPostgres{db: db}
}

func (r *UsuarioPostgres) Create(ctx context.Context, u *models.Usuario) error {
	query := `INSERT INTO usuarios (id, nome, email, senha_hash, cargo, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	now := time.Now()
	_, err := r.db.ExecContext(ctx, query, u.ID, u.Nome, u.Email, u.Senha, u.Cargo, now, now)
	return err
}

func (r *UsuarioPostgres) GetByEmail(ctx context.Context, email string) (*models.Usuario, error) {
	query := `SELECT id, nome, email, senha_hash, cargo, created_at, updated_at, deleted_at 
              FROM usuarios WHERE email = $1 AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, email)

	var u models.Usuario
	err := row.Scan(&u.ID, &u.Nome, &u.Email, &u.Senha, &u.Cargo, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UsuarioPostgres) GetByID(ctx context.Context, id uuid.UUID) (*models.Usuario, error) {
	query := `SELECT id, nome, email, senha_hash, cargo, created_at, updated_at, deleted_at 
              FROM usuarios WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRowContext(ctx, query, id)

	var u models.Usuario
	err := row.Scan(&u.ID, &u.Nome, &u.Email, &u.Senha, &u.Cargo, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UsuarioPostgres) GetByCargo(ctx context.Context, cargo string) ([]*models.Usuario, error) {
	query := `SELECT id, nome, email, cargo, created_at, updated_at, deleted_at 
              FROM usuarios WHERE cargo = $1 AND deleted_at IS NULL ORDER BY nome ASC`
	rows, err := r.db.QueryContext(ctx, query, cargo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []*models.Usuario
	for rows.Next() {
		var u models.Usuario
		if err := rows.Scan(&u.ID, &u.Nome, &u.Email, &u.Cargo, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, &u)
	}
	return usuarios, nil
}

func (r *UsuarioPostgres) Update(ctx context.Context, u *models.Usuario) error {
	query := `UPDATE usuarios SET nome = $1, email = $2, cargo = $3, updated_at = $4 
              WHERE id = $5 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query, u.Nome, u.Email, u.Cargo, time.Now(), u.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("usuário não encontrado")
	}
	return nil
}

func (r *UsuarioPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE usuarios SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL`
	res, err := r.db.ExecContext(ctx, query, time.Now(), id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("usuário não encontrado")
	}
	return nil
}
