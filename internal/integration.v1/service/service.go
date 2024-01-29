package service

import (
	"context"
	"errors"
	"fmt"

	"integration.v1/internal/domain"
)

var ErrGameNotFound = errors.New("game not found")

type service struct {
	repo Repository
	cl   Client
}

func New(c Client, r Repository) *service {
	return &service{
		repo: r,
		cl:   c,
	}
}

func (s *service) GetBalance(ctx context.Context, playerID int) (*domain.Player, error) {
	// Получение информации об игроке по ID
	player, err := s.repo.GetPlayer(playerID)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (s *service) GetGameByID(ctx context.Context, platform string, gameID int) (domain.Game, error) {
	//Получаем список игр
	games, err := s.cl.GetGames(ctx, platform)
	if err != nil {
		return domain.Game{}, err
	}
	var g domain.Game
	//Выбираем игру по ID
	for _, game := range games {
		if gameID == int(game.ID) {
			g = game
			return g, nil
		}
	}
	return domain.Game{}, ErrGameNotFound
}

func (s *service) SendBet(ctx context.Context, playerID int, amount int32) (int32, error) {
	// Получение информации об игроке по ID
	player, err := s.repo.GetPlayer(playerID)
	if err != nil {
		return 0, err
	}

	// Проверка достаточности баланса
	if player.Balance < amount {
		return 0, fmt.Errorf("low balance")
	}

	// Вычитание суммы из баланса и обновление в памяти
	newBalance := player.Balance - amount
	err = s.repo.UpdateBalance(playerID, newBalance)
	if err != nil {
		return 0, err
	}
	return newBalance, nil
}
