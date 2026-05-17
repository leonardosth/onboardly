package models

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	ID        uuid.UUID  `json:"id"`
	Nome      string     `json:"nome" validate:"required,min=3"`
	Email     string     `json:"email" validate:"required,email"`
	Senha     string     `json:"-"`
	Cargo     string     `json:"cargo" validate:"required,oneof=Admin Analista"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required"`
}

type RegisterRequest struct {
	Nome  string `json:"nome" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required,min=6"`
	Cargo string `json:"cargo" validate:"required,oneof=Admin Analista"`
}

type AuthResponse struct {
	Token string  `json:"token"`
	User  Usuario `json:"user"`
}
