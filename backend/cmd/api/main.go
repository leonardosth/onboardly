package main

import (
	"context"
	"log/slog"
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
	// Inicializa Logger Estruturado (slog)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("Arquivo .env não encontrado, utilizando variáveis de sistema")
	}

	db, err := database.Connect()
	if err != nil {
		slog.Error("Falha crítica ao conectar no banco", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	slog.Info("Conectado ao PostgreSQL com sucesso.")

	// Clientes
	clientRepo := repository.NewClientPostgres(db)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handlers.NewClientHandler(clientService)

	// Analistas
	analistaRepo := repository.NewAnalistaPostgres(db)
	analistaService := service.NewAnalistaService(analistaRepo)
	analistaHandler := handlers.NewAnalistaHandler(analistaService)

	// Projetos
	projetoRepo := repository.NewProjetoPostgres(db)
	projetoService := service.NewProjetoService(projetoRepo)
	projetoHandler := handlers.NewProjetoHandler(projetoService)

	// Reuniões
	reuniaoRepo := repository.NewReuniaoPostgres(db)
	reuniaoService := service.NewReuniaoService(reuniaoRepo)
	reuniaoHandler := handlers.NewReuniaoHandler(reuniaoService)

	mux := http.NewServeMux()

	// Endpoints de Clientes
	mux.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			clientHandler.Create(w, r)
		case http.MethodGet:
			clientHandler.GetAll(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/clientes/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			clientHandler.GetByID(w, r)
		case http.MethodPut:
			clientHandler.Update(w, r)
		case http.MethodDelete:
			clientHandler.Delete(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Endpoints de Analistas
	mux.HandleFunc("/analistas", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			analistaHandler.Create(w, r)
		case http.MethodGet:
			analistaHandler.GetAll(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/analistas/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			analistaHandler.GetByID(w, r)
		case http.MethodPut:
			analistaHandler.Update(w, r)
		case http.MethodDelete:
			analistaHandler.Delete(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Endpoints de Projetos
	mux.HandleFunc("/projetos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			projetoHandler.Create(w, r)
		case http.MethodGet:
			projetoHandler.GetAll(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/projetos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			projetoHandler.GetByID(w, r)
		case http.MethodPut:
			projetoHandler.Update(w, r)
		case http.MethodDelete:
			projetoHandler.Delete(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint Dashboard
	mux.HandleFunc("/dashboard/stats", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		projetoHandler.GetStats(w, r)
	})

	// Endpoints de Reuniões
	mux.HandleFunc("/reunioes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			reuniaoHandler.Create(w, r)
		case http.MethodGet:
			reuniaoHandler.GetAll(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/reunioes/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			reuniaoHandler.GetByID(w, r)
		case http.MethodPut:
			reuniaoHandler.Update(w, r)
		case http.MethodDelete:
			reuniaoHandler.Delete(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Aplicar Middleware de CORS
	handler := corsMiddleware(mux)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info("Servidor rodando na porta 8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Erro no servidor", "error", err)
			os.Exit(1)
		}
	}()

	<-stop
	slog.Info("Desligando o servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Erro no Graceful Shutdown", "error", err)
		os.Exit(1)
	}
	slog.Info("Servidor finalizado com segurança.")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
