package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nomannaq/simpledbgolang/internal/database"
)

// SetHandler handles the HTTP POST request to set a key-value pair
func SetHandler(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		db.Set(payload.Key, payload.Value)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Key '%s' set successfully\n", payload.Key)
	}
}

// GetHandler handles the HTTP GET request to retrieve a value by key
func GetHandler(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Missing 'key' parameter", http.StatusBadRequest)
			return
		}
		value, err := db.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{key: value})
	}
}

// DeleteHandler handles the HTTP DELETE request to delete a key-value pair
func DeleteHandler(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Missing 'key' parameter", http.StatusBadRequest)
			return
		}
		db.Delete(key)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Key '%s' deleted successfully\n", key)
	}
}
