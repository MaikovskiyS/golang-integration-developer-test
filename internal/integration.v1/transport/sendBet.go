package transport

import (
	"context"
	"fmt"

	"integration.v1/gen/gamer"
)

const transportSendBetErr string = "Transport-SendBet"

// SendBet реализует метод SendBet.
func (s *integrationServer) SendBet(ctx context.Context, req *gamer.SendBetRequest) (*gamer.SendBetResponse, error) {

	request, err := NewSendBetRequest(req)
	if err != nil {
		return &gamer.SendBetResponse{}, fmt.Errorf("%s: %w", transportSendBetErr, err)
	}
	err = request.Validate()
	if err != nil {
		return &gamer.SendBetResponse{}, fmt.Errorf("%s: %w", transportSendBetErr, err)
	}

	newBalance, err := s.svc.SendBet(ctx, request.Player.Id, request.Amount)
	if err != nil {
		return &gamer.SendBetResponse{}, err
	}

	response := &gamer.SendBetResponse{
		Balance: newBalance,
	}

	return response, nil
}
