package tenant

import (
	"encoding/json"
	"os"
	"time"
)

type Bill struct {
	Tenant string `json:"tenant"`
	Usage  int    `json:"usage"`
	Time   string `json:"time"`
}

func Export(id string) {

	b := Bill{
		Tenant: id,
		Usage:  Get(id),
		Time:   time.Now().UTC().Format(time.RFC3339),
	}

	f, _ := os.OpenFile(
		"billing.jsonl",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	defer f.Close()

	j, _ := json.Marshal(b)

	f.Write(append(j, '\n'))
}
