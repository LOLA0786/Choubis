package shield

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

// Simple rule engine (v1)
var dangerousPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)bypass.*security`),
	regexp.MustCompile(`(?i)ignore.*rules`),
	regexp.MustCompile(`(?i)jailbreak`),
	regexp.MustCompile(`(?i)hack`),
}

// Extract user prompt
func InspectRequest(r *http.Request) (bool, string) {

	if r.Method != "POST" {
		return false, ""
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false, ""
	}

	// Restore body
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		return false, ""
	}

	msgs, ok := payload["messages"].([]interface{})
	if !ok {
		return false, ""
	}

	for _, m := range msgs {

		msg := m.(map[string]interface{})
		content, _ := msg["content"].(string)

		for _, re := range dangerousPatterns {
			if re.MatchString(content) {
				return true, content
			}
		}
	}

	return false, ""
}

// Simple risk score (0-100)
func RiskScore(text string) int {

	score := 0

	for _, re := range dangerousPatterns {
		if re.MatchString(text) {
			score += 30
		}
	}

	if len(text) > 500 {
		score += 10
	}

	if score > 100 {
		score = 100
	}

	return score
}

// RiskScore returns a simple 0-100 risk score
