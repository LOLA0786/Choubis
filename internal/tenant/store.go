package tenant

import (
	"encoding/json"
	"os"
	"sync"
)

type Tenant struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Quota int    `json:"quota"`
}

var mu sync.Mutex
var dbFile = "tenants.json"

func LoadAll() ([]Tenant, error) {

	b, err := os.ReadFile(dbFile)
	if err != nil {
		return []Tenant{}, nil
	}

	var t []Tenant
	json.Unmarshal(b, &t)

	return t, nil
}

func Save(t []Tenant) error {

	b, _ := json.MarshalIndent(t, "", "  ")

	return os.WriteFile(dbFile, b, 0644)
}

func AddTenant(nt Tenant) error {

	mu.Lock()
	defer mu.Unlock()

	all, _ := LoadAll()

	all = append(all, nt)

	return Save(all)
}

func GetTenant(id string) *Tenant {

	all, _ := LoadAll()

	for _, t := range all {
		if t.ID == id {
			return &t
		}
	}

	return nil
}
