package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func getIDFromPath(r *http.Request, prefix string) (uuid.UUID, error) {
	path := strings.TrimPrefix(r.URL.Path, prefix)
	if path == "" || strings.Contains(path, "/") {
		return uuid.Nil, errors.New("id inválido")
	}
	return uuid.Parse(path)
}
