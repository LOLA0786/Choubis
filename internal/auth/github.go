package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var conf *oauth2.Config

var allowedUsers map[string]bool
var allowedOrg string

func Init() {

	conf = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:9000/auth/callback",
		Scopes:       []string{"read:user", "read:org"},
		Endpoint:     github.Endpoint,
	}

	allowedUsers = make(map[string]bool)

	users := os.Getenv("GITHUB_ALLOWED_USERS")
	for _, u := range strings.Split(users, ",") {
		if u != "" {
			allowedUsers[strings.TrimSpace(u)] = true
		}
	}

	allowedOrg = os.Getenv("GITHUB_ALLOWED_ORG")
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

	// Get user
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer resp.Body.Close()

	var user map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&user)

	login := user["login"].(string)

	// Check whitelist
	if !isAllowed(login, client) {
		w.WriteHeader(403)
		w.Write([]byte("Access denied"))
		return
	}

	// Session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "choubis_auth",
		Value:    login,
		Path:     "/",
		HttpOnly: true,
	})

	http.Redirect(w, r, "/admin/", 302)
}

func isAllowed(user string, client *http.Client) bool {

	// User whitelist
	if allowedUsers[user] {
		return true
	}

	// Org whitelist
	if allowedOrg == "" {
		return false
	}

	req, _ := http.NewRequest(
		"GET",
		"https://api.github.com/user/orgs",
		nil,
	)

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var orgs []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&orgs)

	for _, o := range orgs {
		if o["login"] == allowedOrg {
			return true
		}
	}

	return false
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
