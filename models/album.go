package models

// Album represents a music album with all its details
type Album struct {
	ID          int
	Title       string
	Artist      string
	CoverURL    string
	ReleaseYear int
	Genre       string
	Description string
	Tracks      []Track
}

// Track represents an individual song in an album
type Track struct {
	Number      int
	Title       string
	Duration    string // Format: "3:45"
	Description string
}

func NewTrack(no int, title, duration string) Track {
	return Track{
		Number:      no,
		Title:       title,
		Duration:    duration,
		Description: "",
	}
}
