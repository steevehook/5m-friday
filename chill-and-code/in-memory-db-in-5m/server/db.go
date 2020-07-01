package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type memoryDB struct {
	items map[string]string
	mu    sync.RWMutex
}

func newDB() memoryDB {
	f, err := os.Open("db.json")
	if err != nil {
		return memoryDB{items: map[string]string{}}
	}
	items := map[string]string{}
	if err := json.NewDecoder(f).Decode(&items); err != nil {
		fmt.Println("could not decode", err.Error())
		return memoryDB{items: map[string]string{}}
	}
	return memoryDB{items: items}
}

func (m memoryDB) save() {
	f, err := os.Create("db.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(f).Encode(m.items); err != nil {
		log.Fatal(err)
	}
}

func (m memoryDB) set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[key] = value
}

func (m memoryDB) get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, found := m.items[key]
	return value, found
}

func (m memoryDB) delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, key)
}
