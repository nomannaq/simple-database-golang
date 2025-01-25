package database

import (
	"errors"
	"sync"
)

// Database represents a simple in-memory key-value store
type Database struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewDatabase creates a new instance of the in-memory database
func NewDatabase() *Database {
	return &Database{
		store: make(map[string]string),
	}
}

// Set adds or updates a key-value pair in the database
func (db *Database) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.store[key] = value
}

// Get retrieves the value for a given key from the database
func (db *Database) Get(key string) (string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, exists := db.store[key]
	if !exists {
		return "", errors.New("key not found")
	}
	return value, nil
}

// Delete removes a key-value pair from the database
func (db *Database) Delete(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.store, key)
}
