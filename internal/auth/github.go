package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var conf *oauth2.Config

func Init() {

	conf = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:9000/auth/callback",
		Scopes:       []string{"read:user"},
		Endpoint:     github.Endpoint,
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)

	http.Redirect(w, r, url, http.StatusFound)
}

func Callback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")

	tok, err := conf.Exchange(context.Background(), code)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	client := conf.Client(context.Background(), tok)

	resp, _ := client.Get("https://api.github.com/user")

	defer resp.Body.Close()

	var user map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&user)

	// Simple session cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "choubis_auth",
		Value: user["login"].(string),
		Path:  "/",
	})

	http.Redirect(w, r, "/admin/", 302)
}

func Check(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := r.Cookie("choubis_auth")

		if err != nil {
			http.Redirect(w, r, "/auth/login", 302)
			return
		}

		next.ServeHTTP(w, r)
	})
}
