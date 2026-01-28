package vault

import (
	"log"
	"strings"
)

type Enforcer struct {
	TEEActive      bool
	MinSafetyScore int
}

func (e *Enforcer) EvaluateSafety(reasoning string) bool {
	// Simple yet complicated logic: Detect "Weak Reasoning"
	score := 100
	riskWords := []string{"uncertain", "maybe", "hypothetically", "bypass"}
	
	for _, word := range riskWords {
		if strings.Contains(strings.ToLower(reasoning), word) {
			score -= 30
		}
	}

	if score < e.MinSafetyScore {
		log.Printf("🛑 [KILL-SWITCH] Safety Score %d below threshold %d. Terminating connection.", score, e.MinSafetyScore)
		return false
	}
	return true
}
