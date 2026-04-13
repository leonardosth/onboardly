package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/leonardosth/onboardly/internal/database"
	"github.com/leonardosth/onboardly/internal/handlers"
	"github.com/leonardosth/onboardly/internal/repository"
	"github.com/leonardosth/onboardly/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, utilizando variáveis de sistema")
	}

	// 2. Conectar ao Banco de Dados
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Falha crítica ao conectar no banco: %v", err)
	}
	defer db.Close()
	log.Println("Conectado ao PostgreSQL com sucesso.")

	// 3. Injeção de Dependências (Wiring)
	clientRepo := repository.NewClientPostgres(db)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handlers.NewClientHandler(clientService)

	// 4. Rotas (Utilizando net/http padrão do Go 1.22+)
	mux := http.NewServeMux()

	// Endpoints de Clientes
	mux.HandleFunc("POST /clientes", clientHandler.Create)
	mux.HandleFunc("GET /clientes", clientHandler.GetAll)
	mux.HandleFunc("GET /clientes/{id}", clientHandler.GetByID)
	mux.HandleFunc("PUT /clientes/{id}", clientHandler.Update)
	mux.HandleFunc("DELETE /clientes/{id}", clientHandler.Delete)

	// 5. Configurar o Servidor com Graceful Shutdown
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Canal para interceptar sinal de parada do sistema operacional (Ctrl+C, Docker stop)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Servidor rodando na porta 8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro no servidor: %v", err)
		}
	}()

	<-stop // Bloqueia até recebermos um sinal
	log.Println("Desligando o servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Erro no Graceful Shutdown: %v", err)
	}
	log.Println("Servidor finalizado com segurança.")
}
