package models

import "time"

// Lane represents info about a single shooting lane
type Lane struct {
	ID        int       `json:"id"`                // Lane ID or lane number
	Shooter   string    `json:"shooter,omitempty"` // Optional name(s) of shooter(s)
	Rental    bool      `json:"rental,omitempty"`  // Whether they have a rental
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Status    string    `json:"status,omitempty"` // e.g. "green", "yellow", "red", "maintenance"
}
