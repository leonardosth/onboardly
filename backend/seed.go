package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/leonardosth/onboardly/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	godotenv.Load(".env")
	godotenv.Load("backend/.env")

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	// 1. Criar Admin
	createUsuario(db, "Leonardo Admin", "leonardo@onboardly.com", "admin123", "Admin")

	// 2. Criar Analista de Teste (ID Fixo para facilitar o frontend/testes se necessário, ou apenas garantir existência)
	// Usando o ID que o frontend está tentando usar (visto no log anterior)
	analistaID, _ := uuid.Parse("418d9f5d-2304-4c56-aa7b-9065c99848fd")
	createUsuarioWithID(db, analistaID, "Leonardo Lima", "leonardo.lima@onboarly.com", "analista123", "Analista")

	// 3. Criar Cliente de Teste
	clienteID, _ := uuid.Parse("1b57536c-6fd8-40e7-ad57-ddf8fa665b1e")
	createCliente(db, clienteID, "Conta Azul", "05206246000138")

	fmt.Println("\n✅ Seed finalizado com sucesso!")
}

func createUsuario(db *sql.DB, nome, email, senha, cargo string) {
	createUsuarioWithID(db, uuid.New(), nome, email, senha, cargo)
}

func createUsuarioWithID(db *sql.DB, id uuid.UUID, nome, email, senha, cargo string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(senha), 10)
	query := `
		INSERT INTO usuarios (id, nome, email, senha_hash, cargo, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		ON CONFLICT (email) DO UPDATE SET nome = $2, cargo = $5, updated_at = CURRENT_TIMESTAMP
	`
	_, err := db.Exec(query, id, nome, email, string(hash), cargo)
	if err != nil {
		log.Printf("❌ Erro ao criar usuário %s: %v", email, err)
	} else {
		fmt.Printf("👤 Usuário [%s] pronto (ID: %s)\n", cargo, id)
	}
}

func createCliente(db *sql.DB, id uuid.UUID, nome, cnpj string) {
	query := `
		INSERT INTO clientes (id, nome, cnpj, created_at, updated_at) 
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		ON CONFLICT (cnpj) DO UPDATE SET nome = $2, updated_at = CURRENT_TIMESTAMP
	`
	_, err := db.Exec(query, id, nome, cnpj)
	if err != nil {
		log.Printf("❌ Erro ao criar cliente %s: %v", nome, err)
	} else {
		fmt.Printf("🏢 Cliente [%s] pronto (ID: %s)\n", nome, id)
	}
}
