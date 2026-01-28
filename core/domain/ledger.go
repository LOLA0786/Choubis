package domain

import "time"

// LedgerEntry represents a single "Thought Trace" from an AI Agent
type LedgerEntry struct {
    AgentID     string    `json:"agent_id"`
    IntentID    string    `json:"intent_id"`
    Timestamp   time.Time `json:"timestamp"`
    Reasoning   string    `json:"reasoning"` // The "Chain of Thought"
    Action      string    `json:"action"`    // The actual API call
    Hash        string    `json:"hash"`      // Signed hash for Insurance/Audit
}

type LedgerRepository interface {
    Store(entry LedgerEntry) error
    Verify(intentID string) (bool, error)
}
