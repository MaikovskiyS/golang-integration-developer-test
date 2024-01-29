package domain

import (
	"strconv"
)

type GameID int

func (g GameID) String() string {
	return strconv.Itoa(int(g))
}

// Game структура для хранения данных об играх.
type Game struct {
	ID                     GameID `json:"id"`
	Title                  string `json:"title"`
	Thumbnail              string `json:"thumbnail"`
	Description            string `json:"short_description"`
	Game_url               string `json:"game_url"`
	Genre                  string `json:"genre"`
	Platform               string `json:"platform"`
	Publisher              string `json:"publisher"`
	Developer              string `json:"developer"`
	Release_date           string `json:"release_date"`
	Freetogame_profile_url string `json:"freetogame_profile_url"`
}
