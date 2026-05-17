package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/leonardosth/onboardly/internal/service"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthMiddleware(authService service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "autorização necessária"})
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "formato de token inválido"})
				return
			}

			tokenString := parts[1]
			claims, err := authService.ValidateToken(tokenString)
			if err != nil {
				respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "token inválido ou expirado"})
				return
			}

			// Adicionar claims ao contexto se necessário
			ctx := context.WithValue(r.Context(), UserContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := r.Context().Value(UserContextKey)
			if val == nil {
				respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "usuário não autenticado"})
				return
			}

			claims, ok := val.(*jwt.MapClaims)
			if !ok {
				respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "erro ao processar permissões"})
				return
			}

			userRole, ok := (*claims)["cargo"].(string)
			if !ok {
				respondJSON(w, http.StatusForbidden, map[string]string{"error": "permissão negada: cargo não identificado"})
				return
			}

			authorized := false
			for _, role := range roles {
				if userRole == role {
					authorized = true
					break
				}
			}

			if !authorized {
				respondJSON(w, http.StatusForbidden, map[string]string{"error": "permissão negada para este recurso"})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("PANIC RECOVERED", "error", err)
				respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "erro interno do servidor"})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
