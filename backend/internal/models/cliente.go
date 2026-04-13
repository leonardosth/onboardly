package models

import (
	"time"

	"github.com/google/uuid"
)

type Cliente struct {
	ID           uuid.UUID `json:"id" db:"id"`
	NomeFantasia string    `json:"nome" db:"nome_fantasia"`
	CNPJ         string    `json:"cnpj" db:"cnpj"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
