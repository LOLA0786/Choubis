package shield

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var auditMu sync.Mutex

type AuditEvent struct {
	Timestamp string `json:"ts"`
	Path      string `json:"path"`
	Risk       int    `json:"risk"`
	Action     string `json:"action"`
	Reason     string `json:"reason"`
}

func auditPath() string {

	exe, err := os.Executable()
	if err != nil {
		fmt.Println("AUDIT exe error:", err)
		return "logs/events.jsonl"
	}

	fmt.Println("AUDIT exe:", exe)

	base := filepath.Dir(filepath.Dir(exe))

	fmt.Println("AUDIT base:", base)

	return filepath.Join(base, "logs", "events.jsonl")
}

func LogEvent(e AuditEvent) {

	auditMu.Lock()
	defer auditMu.Unlock()

	path := auditPath()

	fmt.Println("AUDIT path:", path)

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		fmt.Println("AUDIT mkdir error:", err)
		return
	}

	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Println("AUDIT open error:", err)
		return
	}
	defer f.Close()

	e.Timestamp = time.Now().UTC().Format(time.RFC3339)

	b, _ := json.Marshal(e)

	_, err = f.Write(append(b, '\n'))
	if err != nil {
		fmt.Println("AUDIT write error:", err)
	}
}
