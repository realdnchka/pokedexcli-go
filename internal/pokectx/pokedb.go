package pokectx

import (
	"sync"
)

type databaseCreatable interface {
	Create()
}
type databaseWR interface {
	Write(any)
	Read(any)
}

type DB struct {
	databaseCreatable
	databaseWR
	entry map[string]Pokemon
	mu sync.RWMutex
}

func (db *DB) Create() *DB {
	db.mu.Lock()
	db.entry = make(map[string]Pokemon, 0)
	db.mu.Unlock()
	return db
}

func (db *DB) Write(key string, value Pokemon) *DB {
	db.mu.RLock()
	db.entry[key] = value
	db.mu.RUnlock()
	return db
}

func (db *DB) Read(key string) (Pokemon, bool) {
	if _, ok := db.entry[key]; !ok {
		return Pokemon{}, false
	}
	return db.entry[key], true
}

func (db *DB) ReadAll() (poks []Pokemon) {
	for _, v := range db.entry {
		poks = append(poks, v)
	}
	return poks
}