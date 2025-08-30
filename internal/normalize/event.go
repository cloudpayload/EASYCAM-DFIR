package normalize

import "time"

// NormalizedEvent represents a standardized event structure
// that can be used across different DFIR tools and platforms
type NormalizedEvent struct {
	Timestamp   time.Time `json:"timestamp"`
	Actor       string    `json:"actor"`
	Action      string    `json:"action"`
	Target      string    `json:"target"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	RequestID   string    `json:"request_id"`
	Status      string    `json:"status"`
	Provider    string    `json:"provider"`
	SourcePath  string    `json:"source_path"`
	// Add any other common fields
}
