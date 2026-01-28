package shield

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"choubis/pkg/vault"
)

type ShieldProxy struct {
	Proxy    *httputil.ReverseProxy
	Target   *url.URL
	Enforcer *vault.Enforcer
}

func NewShieldProxy(targetURL string) (*ShieldProxy, error) {
	target, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	enforcer := &vault.Enforcer{MinSafetyScore: 50}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{ServerName: "api.openai.com"},
	}

	proxy.Director = func(req *http.Request) {
		req.Host = target.Host
		req.URL.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.Header.Set("Host", target.Host)
		req.Header.Set("User-Agent", "Choubis-Sovereign-Agent/1.0")
		
		if !modifyRequestBody(req, enforcer) {
			// This tells the proxy to stop processing
			req.URL.Path = "/kill-switch-triggered"
		}
	}

	return &ShieldProxy{Proxy: proxy, Target: target, Enforcer: enforcer}, nil
}

func modifyRequestBody(req *http.Request, e *vault.Enforcer) bool {
	if req.Body == nil { return true }

	body, _ := io.ReadAll(req.Body)
	content := string(body)

	// Trigger Kill Switch if 'bypass' is detected in the prompt
	if !e.EvaluateSafety(content) {
		return false
	}

	redactedContent := strings.ReplaceAll(content, "Chandan Galani", "[REDACTED_USER]")
	req.Body = io.NopCloser(bytes.NewBufferString(redactedContent))
	req.ContentLength = int64(len(redactedContent))
	return true
}
