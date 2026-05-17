package service

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leonardosth/onboardly/internal/models"
	"github.com/leonardosth/onboardly/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, req models.RegisterRequest) (*models.Usuario, error)
	Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error)
	ValidateToken(tokenString string) (*jwt.MapClaims, error)
}

type authService struct {
	repo repository.UsuarioRepository
}

func NewAuthService(repo repository.UsuarioRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(ctx context.Context, req models.RegisterRequest) (*models.Usuario, error) {
	// Verificar se usuário já existe
	existing, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("usuário com este email já existe")
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.Usuario{
		ID:    uuid.New(),
		Nome:  req.Nome,
		Email: req.Email,
		Senha: string(hashedPassword),
		Cargo: req.Cargo,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		slog.Error("erro ao buscar usuário no banco", "error", err, "email", req.Email)
		return nil, err
	}
	if user == nil {
		slog.Warn("tentativa de login: usuário não encontrado", "email", req.Email)
		return nil, errors.New("credenciais inválidas")
	}

	// Verificar senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(req.Senha)); err != nil {
		slog.Warn("tentativa de login: senha incorreta", "email", req.Email)
		return nil, errors.New("credenciais inválidas")
	}

	slog.Info("login realizado com sucesso", "usuario_id", user.ID, "email", user.Email)

	// Gerar Token JWT
	token, err := s.generateJWT(user)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *authService) generateJWT(user *models.Usuario) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "onboardly_super_secret_key" // Fallback para desenvolvimento
	}

	claims := jwt.MapClaims{
		"sub":   user.ID.String(),
		"email": user.Email,
		"nome":  user.Nome,
		"cargo": user.Cargo,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (s *authService) ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "onboardly_super_secret_key"
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inesperado")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("token inválido")
}
