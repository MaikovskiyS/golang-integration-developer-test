package transport

import (
	"strconv"

	"integration.v1/gen/gamer"
)

type GetBalanceRequest struct {
	Token    string
	Player   ReqPlayer
	Platform string
	Currency ReqCurrency
	GameID   int
}

type SendBetRequest struct {
	Token         string
	Player        ReqPlayer
	Platform      string
	Currency      ReqCurrency
	GameID        int
	TransactionID int
	Amount        int32
}

func NewGetBalanceRequest(req *gamer.GetBalanceRequest) (*GetBalanceRequest, error) {
	if req.String() == "" {
		return &GetBalanceRequest{}, ErrInvalidParams
	}
	gameID, err := strconv.Atoi(req.General.GameId)
	if err != nil {
		return &GetBalanceRequest{}, ErrInvalidGameID
	}
	playerID, err := strconv.Atoi(req.General.Player.Id)
	if err != nil {
		return &GetBalanceRequest{}, ErrInvalidPlayerParams
	}

	return &GetBalanceRequest{
		Token:    req.General.Token,
		Player:   ReqPlayer{Id: playerID, NickName: req.General.Player.Nickname},
		GameID:   gameID,
		Platform: req.General.Platform,
		Currency: ReqCurrency{Code: req.General.Currency.Code, Name: req.General.Currency.Name},
	}, nil
}
func NewSendBetRequest(req *gamer.SendBetRequest) (*SendBetRequest, error) {
	if req.String() == "" {
		return &SendBetRequest{}, ErrInvalidParams
	}
	gameID, err := strconv.Atoi(req.General.GameId)
	if err != nil {
		return &SendBetRequest{}, ErrInvalidGameID
	}
	playerID, err := strconv.Atoi(req.General.Player.Id)
	if err != nil {
		return &SendBetRequest{}, ErrInvalidPlayerParams
	}
	tr_id, err := strconv.Atoi(req.General.GameId)
	if err != nil {
		return &SendBetRequest{}, ErrInvalidTransactionID
	}
	return &SendBetRequest{
		Token:         req.General.Token,
		Player:        ReqPlayer{Id: playerID, NickName: req.General.Player.Nickname},
		GameID:        gameID,
		Platform:      req.General.Platform,
		Currency:      ReqCurrency{Code: req.General.Currency.Code, Name: req.General.Currency.Name},
		TransactionID: tr_id,
		Amount:        req.Amount,
	}, nil
}

type ReqPlayer struct {
	Id       int
	NickName string
}
type ReqCurrency struct {
	Code string
	Name string
}
