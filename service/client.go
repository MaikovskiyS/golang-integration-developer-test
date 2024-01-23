package service

// Games list from rest API
type Games []struct {
	// todo: parse output from https://www.freetogame.com/api/games?platform=pc
	// Credits: freetogame.com
}

// Client http service interface
type Client interface {
	GetGames(platform string) (Games, error)
	GetBalance(playerID string) (int64, error)
}
