package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/leonardosth/onboardly/internal/models"
	"github.com/leonardosth/onboardly/internal/service"
	"golang.org/x/crypto/bcrypt"
)

// Mock UsuarioRepository
type mockUsuarioRepository struct {
	mockGetByEmail func(email string) (*models.Usuario, error)
	mockCreate     func(u *models.Usuario) error
	mockGetByID    func(id uuid.UUID) (*models.Usuario, error)
	mockGetByCargo func(cargo string) ([]*models.Usuario, error)
	mockUpdate     func(u *models.Usuario) error
	mockDelete     func(id uuid.UUID) error
}

func (m *mockUsuarioRepository) GetByEmail(ctx context.Context, email string) (*models.Usuario, error) {
	return m.mockGetByEmail(email)
}
func (m *mockUsuarioRepository) Create(ctx context.Context, u *models.Usuario) error {
	return m.mockCreate(u)
}
func (m *mockUsuarioRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Usuario, error) {
	return m.mockGetByID(id)
}
func (m *mockUsuarioRepository) GetByCargo(ctx context.Context, cargo string) ([]*models.Usuario, error) {
	return m.mockGetByCargo(cargo)
}
func (m *mockUsuarioRepository) Update(ctx context.Context, u *models.Usuario) error {
	return m.mockUpdate(u)
}
func (m *mockUsuarioRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return m.mockDelete(id)
}

func TestAuthFlow(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	testUser := &models.Usuario{
		ID:    uuid.New(),
		Nome:  "Test User",
		Email: "test@example.com",
		Senha: string(hashedPassword),
		Cargo: "Analista",
	}

	mockRepo := &mockUsuarioRepository{
		mockGetByEmail: func(email string) (*models.Usuario, error) {
			if email == testUser.Email {
				return testUser, nil
			}
			return nil, nil
		},
	}

	authService := service.NewAuthService(mockRepo)
	authHandler := NewAuthHandler(authService)

	t.Run("Login com sucesso", func(t *testing.T) {
		loginReq := models.LoginRequest{
			Email: "test@example.com",
			Senha: "password123",
		}
		body, _ := json.Marshal(loginReq)
		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		authHandler.Login(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rr.Code)
		}

		respBytes := rr.Body.Bytes()
		var resp models.AuthResponse
		if err := json.Unmarshal(respBytes, &resp); err != nil {
			t.Fatalf("Erro ao decodificar resposta: %v", err)
		}

		if resp.Token == "" {
			t.Error("Token não deveria estar vazio")
		}
		if resp.User.Email != testUser.Email {
			t.Errorf("Esperava email %s, recebeu %s", testUser.Email, resp.User.Email)
		}
		// Verificar que a senha não foi exposta no JSON
		var rawResp map[string]interface{}
		json.Unmarshal(respBytes, &rawResp)
		userObj, ok := rawResp["user"].(map[string]interface{})
		if !ok {
			t.Fatalf("Resposta JSON não contém objeto 'user' válido: %s", string(respBytes))
		}
		if _, exists := userObj["senha"]; exists {
			t.Error("Senha não deve ser exposta na resposta JSON")
		}
	})

	t.Run("Login com senha errada", func(t *testing.T) {
		loginReq := models.LoginRequest{
			Email: "test@example.com",
			Senha: "wrongpassword",
		}
		body, _ := json.Marshal(loginReq)
		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		authHandler.Login(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Esperava status 401, recebeu %d", rr.Code)
		}
	})

	t.Run("Login com email inexistente", func(t *testing.T) {
		loginReq := models.LoginRequest{
			Email: "nonexistent@example.com",
			Senha: "password123",
		}
		body, _ := json.Marshal(loginReq)
		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		authHandler.Login(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Esperava status 401, recebeu %d", rr.Code)
		}
	})
}

func TestAuthMiddleware(t *testing.T) {
	authService := service.NewAuthService(&mockUsuarioRepository{})
	middleware := AuthMiddleware(authService)

	// Handler Dummy para testar se o middleware passa a requisição
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	protectedHandler := middleware(dummyHandler)

	t.Run("Acesso sem token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/protected", nil)
		rr := httptest.NewRecorder()

		protectedHandler.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Esperava status 401, recebeu %d", rr.Code)
		}
	})

	t.Run("Acesso com formato de token inválido", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Invalido token123")
		rr := httptest.NewRecorder()

		protectedHandler.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Esperava status 401, recebeu %d", rr.Code)
		}
	})

	t.Run("Acesso com token válido", func(t *testing.T) {
		// Gerar um token real para o teste
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := &models.Usuario{
			ID:    uuid.New(),
			Nome:  "Test User",
			Email: "test@example.com",
			Senha: string(hashedPassword),
		}
		
		mockRepo := &mockUsuarioRepository{
			mockGetByEmail: func(email string) (*models.Usuario, error) {
				return user, nil
			},
		}
		authServiceWithRepo := service.NewAuthService(mockRepo)
		resp, _ := authServiceWithRepo.Login(context.Background(), models.LoginRequest{
			Email: user.Email,
			Senha: "password123",
		})

		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+resp.Token)
		rr := httptest.NewRecorder()

		// Recriar o middleware com o service que tem o repo mockado (embora ValidateToken não use repo, é bom manter consistência)
		middlewareWithToken := AuthMiddleware(authServiceWithRepo)
		middlewareWithToken(dummyHandler).ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rr.Code)
		}
	})
}
