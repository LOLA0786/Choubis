package admin

import (
	"bufio"
	"net/http"
	"os"
)

func serveAudit(w http.ResponseWriter) {

	f, err := os.Open("logs/events.jsonl")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("["))

	first := true

	for scanner.Scan() {

		if !first {
			w.Write([]byte(","))
		}
		first = false

		w.Write(scanner.Bytes())
	}

	w.Write([]byte("]"))
}
