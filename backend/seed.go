package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/leonardosth/onboardly/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Tenta carregar o .env da pasta atual ou da pasta pai
	godotenv.Load(".env")
	godotenv.Load("backend/.env")

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar no banco: %v. Verifique se as variáveis no .env estão corretas.", err)
	}
	defer db.Close()

	nome := "Leonardo Admin"
	email := "leonardo@onboardly.com"
	senha := "admin123"
	cargo := "CTO"

	// Gera o hash da senha usando bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		log.Fatalf("❌ Erro ao gerar hash da senha: %v", err)
	}

	// SQL para inserir o usuário. 
	// O 'ON CONFLICT' evita erro se você rodar o script mais de uma vez.
	query := `
		INSERT INTO usuarios (id, nome, email, senha_hash, cargo, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		ON CONFLICT (email) DO NOTHING
	`
	
	id := uuid.New()
	result, err := db.Exec(query, id, nome, email, string(hash), cargo)
	if err != nil {
		log.Fatalf("❌ Erro ao executar INSERT: %v", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		fmt.Println("⚠️  Atenção: O usuário já existia no banco de dados.")
	} else {
		fmt.Println("✅ Usuário administrador criado com sucesso!")
		fmt.Printf("🆔 ID: %s\n", id)
	}

	fmt.Println("-------------------------------------------")
	fmt.Printf("📧 Email: %s\n", email)
	fmt.Printf("🔑 Senha: %s\n", senha)
	fmt.Println("-------------------------------------------")
	fmt.Println("Agora você já pode logar no Onboardly!")
}
