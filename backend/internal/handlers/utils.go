package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func validatePayload(w http.ResponseWriter, payload interface{}) bool {
	if err := validate.Struct(payload); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field() + " is " + err.Tag())
		}
		respondJSON(w, http.StatusBadRequest, map[string]interface{}{"errors": errors})
		return false
	}
	return true
}

func getIDFromPath(r *http.Request, prefix string) (uuid.UUID, error) {
	path := strings.TrimPrefix(r.URL.Path, prefix)
	if path == "" || strings.Contains(path, "/") {
		return uuid.Nil, errors.New("id inválido")
	}
	return uuid.Parse(path)
}
