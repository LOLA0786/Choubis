package main

import (
    "log"
    "net/http"
    "choubis/internal/shield"
    "github.com/fatih/color"
)

func main() {
    color.Cyan("🛡️ Choubis (चौबीस) Sovereign AI Gateway Starting...")
    
    // Target: Proxying to OpenAI as a test (Shadow AI capture)
    proxy, err := shield.NewShieldProxy("https://api.openai.com")
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("[CHOUBIS] Routing request: %s %s", r.Method, r.URL.Path)
        proxy.Proxy.ServeHTTP(w, r)
    })

    color.Green("⚡ Proxy live on :9000 | Wrapping all outbound AI traffic.")
    log.Fatal(http.ListenAndServe(":9000", nil))
}
