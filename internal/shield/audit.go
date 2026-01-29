package shield

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var auditMu sync.Mutex

type AuditEvent struct {
	Timestamp string `json:"ts"`
	Tenant    string `json:"tenant"`
	Path      string `json:"path"`
	Risk       int    `json:"risk"`
	Action     string `json:"action"`
	Reason     string `json:"reason"`
}

func LogEvent(tenant string, e AuditEvent) {

	auditMu.Lock()
	defer auditMu.Unlock()

	path := filepath.Join("logs", tenant+".jsonl")

	os.MkdirAll("logs", 0755)

	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return
	}
	defer f.Close()

	e.Timestamp = time.Now().UTC().Format(time.RFC3339)
	e.Tenant = tenant

	b, _ := json.Marshal(e)

	f.Write(append(b, '\n'))
}
