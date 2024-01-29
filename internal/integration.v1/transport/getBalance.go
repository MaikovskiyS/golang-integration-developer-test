package transport

import (
	"context"
	"fmt"

	"integration.v1/gen/gamer"
)

const transportGetBalanceErr string = "Transport-GetBalance"

// GetBalance реализует метод GetBalance
func (s *integrationServer) GetBalance(ctx context.Context, req *gamer.GetBalanceRequest) (*gamer.GetBalanceResponse, error) {

	request, err := NewGetBalanceRequest(req)
	if err != nil {
		return &gamer.GetBalanceResponse{}, fmt.Errorf("%s: %w", transportGetBalanceErr, err)
	}
	err = request.Validate()
	if err != nil {
		return &gamer.GetBalanceResponse{}, fmt.Errorf("%s: %w", transportGetBalanceErr, err)
	}

	game, err := s.svc.GetGameByID(ctx, request.Platform, request.GameID)
	if err != nil {
		return &gamer.GetBalanceResponse{}, fmt.Errorf("%s: %w", transportGetBalanceErr, err)
	}
	player, err := s.svc.GetBalance(ctx, request.Player.Id)
	if err != nil {
		return &gamer.GetBalanceResponse{}, fmt.Errorf("%s: %w", transportGetBalanceErr, err)
	}

	response := &gamer.GetBalanceResponse{
		Game:    &gamer.Game{Id: game.ID.String(), Title: game.Title, ShortDescription: game.Description},
		Balance: player.Balance,
	}

	return response, nil
}
