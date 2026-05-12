package models

import (
	"time"

	"github.com/google/uuid"
)

type Analista struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Nome      string     `json:"nome" db:"nome"`
	Email     string     `json:"email" db:"email"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
