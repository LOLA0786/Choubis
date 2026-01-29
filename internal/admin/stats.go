package admin

import (
	"net/http"
	"encoding/json"
	"os"
)

type Stats struct {
	Total  int `json:"total"`
	Blocks int `json:"blocks"`
}

func collectStats() Stats {

	f, err := os.Open("logs/events.jsonl")
	if err != nil {
		return Stats{}
	}
	defer f.Close()

	var s Stats

	buf := make([]byte, 4096)

	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			if buf[i] == '\n' {
				s.Total++
			}
		}
	}

	return s
}

func serveStats(w http.ResponseWriter) {

	s := collectStats()

	b, _ := json.Marshal(s)

	w.Write(b)
}
