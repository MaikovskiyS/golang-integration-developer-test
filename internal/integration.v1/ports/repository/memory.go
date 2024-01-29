package repository

import (
	"fmt"
	"sync"

	"integration.v1/internal/domain"
)

// MemoryStorage механизм хранения в памяти для информации об игроках и балансах.
type MemoryStorage struct {
	mu      *sync.Mutex
	players map[int]*domain.Player
}

func NewMemoryStorage() *MemoryStorage {
	players := make(map[int]*domain.Player)
	players[1] = &domain.Player{ID: 1, Balance: 100}
	players[2] = &domain.Player{ID: 2, Balance: 10}
	return &MemoryStorage{
		mu:      &sync.Mutex{},
		players: players,
	}
}

// GetPlayer возвращает информацию об игроке по его идентификатору.
func (s *MemoryStorage) GetPlayer(playerID int) (*domain.Player, error) {
	player, ok := s.players[playerID]
	if !ok {
		return &domain.Player{}, fmt.Errorf("player with ID %d not found", playerID)
	}
	return player, nil
}

// UpdateBalance обновляет баланс игрока.
func (s *MemoryStorage) UpdateBalance(playerID int, newBalance int32) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	player, ok := s.players[playerID]
	if !ok {
		return fmt.Errorf("player with ID %d not found", playerID)
	}
	player.Balance = newBalance
	s.players[playerID] = player
	return nil
}
