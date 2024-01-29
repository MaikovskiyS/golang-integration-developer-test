package service

import (
	"context"

	"integration.v1/internal/domain"
)

// Client интерфейс для взаимодействия с внешним сервисом FreeToGame.
type Client interface {
	GetGames(ctx context.Context, platform string) ([]domain.Game, error)
}

// Repository интерфейс для взаимодействия с хранилищкм
type Repository interface {
	GetPlayer(playerID int) (*domain.Player, error)
	UpdateBalance(playerID int, newBalance int32) error
}
