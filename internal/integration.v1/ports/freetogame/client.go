package freetogame

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"integration.v1/internal/domain"
)

const timeout = time.Second * 5
const baseUrl = "https://www.freetogame.com"

// APIClient реализует интерфейс Client для взаимодействия с API FreeToGame.
type apiClient struct {
	cl      *http.Client
	baseURL string
}

func NewFreeToGameClient() *apiClient {
	return &apiClient{
		cl:      &http.Client{Timeout: timeout},
		baseURL: baseUrl,
	}
}

// GetGames отправляет запрос к API FreeToGame и возвращает список игр.
func (a *apiClient) GetGames(ctx context.Context, platform string) ([]domain.Game, error) {
	url := fmt.Sprintf("%s/api/games?platform=%s", a.baseURL, platform)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("RequestErr: %w", err)
	}
	response, err := a.cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка запроса. Статус код: %d", response.StatusCode)
	}

	var games []domain.Game
	err = json.NewDecoder(response.Body).Decode(&games)
	if err != nil {
		return nil, err
	}

	return games, nil
}
