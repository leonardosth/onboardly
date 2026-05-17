package models

import (
	"time"

	"github.com/google/uuid"
)

type StatusReuniao string

const (
	StatusAgendada  StatusReuniao = "Agendada"
	StatusRealizada StatusReuniao = "Realizada"
	StatusRemarcada StatusReuniao = "Remarcada"
	StatusNoShow    StatusReuniao = "No_Show"
)

type Reuniao struct {
	ID           uuid.UUID     `json:"id" db:"id"`
	ProjetoID    uuid.UUID     `json:"projeto_id" db:"projeto_id"`
	AnalistaID   uuid.UUID     `json:"analista_id" db:"analista_id"`
	DataAgendada time.Time     `json:"data_agendada" db:"data_agendada"`
	Status       StatusReuniao `json:"status" db:"status"`
	Observacoes  string        `json:"observacoes" db:"observacoes"`
	CreatedAt    time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at,omitempty" db:"deleted_at"`
}
