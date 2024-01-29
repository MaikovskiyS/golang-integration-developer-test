package transport

import (
	"context"

	"integration.v1/internal/domain"
)

type Service interface {
	GetGameByID(ctx context.Context, platform string, gameID int) (domain.Game, error)
	GetBalance(ctx context.Context, playerID int) (*domain.Player, error)
	SendBet(ctx context.Context, playerID int, amount int32) (int32, error)
}

type integrationServer struct {
	svc Service
}

func New(s Service) *integrationServer {
	return &integrationServer{svc: s}
}
