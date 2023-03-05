package model

// Contest has Name, ID, Kind, State.
type Contest struct {
	Name  string
	ID    string // like "abc123"
	Kind  string // like "Algorithm", "Heuristics"...
	State string // "permanent", "upcoming", "archive"
	// TODO: StateをEnumに
}
