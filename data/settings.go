package data

import "github.com/Aj4x/music-collection/models"

func Settings() models.Settings {
	return models.Settings{
		Username:         "admin",
		Email:            "test@example.com",
		AvailablePlayers: []string{"mpv", "vlc"},
		DefaultPlayer:    "mpv",
		Notifications:    true,
		Theme:            "dark",
	}
}
