package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// Connect initializes and verifies the connection to the PostgreSQL database.
func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Abrir a conexão
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir banco de dados: %w", err)
	}

	// Boas práticas de connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verificar se a conexão está ativa
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	return db, nil
}
