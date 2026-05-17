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

	// --- Inicialização de Repositories ---
	clientRepo := repository.NewClientPostgres(db)
	projetoRepo := repository.NewProjetoPostgres(db)
	reuniaoRepo := repository.NewReuniaoPostgres(db)
	usuarioRepo := repository.NewUsuarioPostgres(db)

	// --- Inicialização de Services ---
	clientService := service.NewClientService(clientRepo)
	usuarioService := service.NewUsuarioService(usuarioRepo)
	projetoService := service.NewProjetoService(projetoRepo)
	reuniaoService := service.NewReuniaoService(reuniaoRepo)
	authService := service.NewAuthService(usuarioRepo)

	// --- Inicialização de Handlers ---
	clientHandler := handlers.NewClientHandler(clientService)
	usuarioHandler := handlers.NewUsuarioHandler(usuarioService)
	projetoHandler := handlers.NewProjetoHandler(projetoService)
	reuniaoHandler := handlers.NewReuniaoHandler(reuniaoService)
	authHandler := handlers.NewAuthHandler(authService)

	mux := http.NewServeMux()

	// --- Middlewares ---
	authMiddleware := handlers.AuthMiddleware(authService)
	adminOnly := handlers.RequireRole("Admin")

	// --- Rotas Públicas ---
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	// --- Rotas Protegidas (Envolvidas individualmente para evitar conflitos de 404) ---
	
	// Clientes
	mux.Handle("/clientes", authMiddleware(http.HandlerFunc(clientHandler.GetAll)))
	mux.Handle("/clientes/", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: clientHandler.GetByID(w, r)
		case http.MethodPut: clientHandler.Update(w, r)
		case http.MethodDelete: clientHandler.Delete(w, r)
		default: http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Analistas
	mux.Handle("/analistas", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: 
			usuarioHandler.GetAnalistas(w, r)
		case http.MethodPost: 
			adminOnly(http.HandlerFunc(usuarioHandler.CreateAnalista)).ServeHTTP(w, r)
		default: 
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))
	mux.Handle("/analistas/", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: 
			usuarioHandler.GetByID(w, r)
		case http.MethodPut: 
			adminOnly(http.HandlerFunc(usuarioHandler.UpdateAnalista)).ServeHTTP(w, r)
		case http.MethodDelete: 
			adminOnly(http.HandlerFunc(usuarioHandler.Delete)).ServeHTTP(w, r)
		default: 
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Projetos
	mux.Handle("/projetos", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: projetoHandler.GetAll(w, r)
		case http.MethodPost: projetoHandler.Create(w, r)
		default: http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))
	mux.Handle("/projetos/", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: projetoHandler.GetByID(w, r)
		case http.MethodPut: projetoHandler.Update(w, r)
		case http.MethodDelete: projetoHandler.Delete(w, r)
		default: http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Dashboard e Reuniões
	mux.Handle("/dashboard/stats", authMiddleware(http.HandlerFunc(projetoHandler.GetStats)))
	mux.Handle("/reunioes", authMiddleware(http.HandlerFunc(reuniaoHandler.GetAll)))
	mux.Handle("/reunioes/", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: reuniaoHandler.GetByID(w, r)
		case http.MethodPut: reuniaoHandler.Update(w, r)
		case http.MethodDelete: reuniaoHandler.Delete(w, r)
		default: http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Aplicar Middleware de Recovery e CORS globalmente
	handler := handlers.RecoveryMiddleware(corsMiddleware(mux))

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
		origin := os.Getenv("CORS_ALLOWED_ORIGIN")
		if origin == "" {
			origin = "http://localhost:5173"
		}
		
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
