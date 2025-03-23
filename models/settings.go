package models

// Settings holds the configuration options for the music collection application.
type Settings struct {
	Username         string   // The user's name
	Email            string   // The user's email address
	AvailablePlayers []string // A list of available music players
	DefaultPlayer    string   // The selected default music player
	Notifications    bool     // Flag to enable or disable notifications
	Theme            string   // The selected theme (for example: "light" or "dark")
}
