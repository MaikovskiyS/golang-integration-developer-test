package transport

import (
	"errors"
)

var (
	ErrInvalidParams        = errors.New("invalid request params")
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidPlayerParams  = errors.New("invalid player params")
	ErrInvalidGameID        = errors.New("invalid game id")
	ErrInvalidPlatform      = errors.New("invalid platform")
	ErrInvalidCurrency      = errors.New("invalid currency")
	ErrInvalidTransactionID = errors.New("invalid transaction_id")
	ErrInvalidAmount        = errors.New("invalid amount")
)

const (
	PlatformMobile string = "mobile"
	PlatformPC     string = "pc"
)

func (r *GetBalanceRequest) Validate() error {
	if r.Token == "" {
		return ErrInvalidToken
	}
	if r.Player.Id <= 0 || r.Player.NickName == "" {
		return ErrInvalidPlayerParams
	}
	if r.GameID <= 0 {
		return ErrInvalidGameID
	}
	if r.Platform != PlatformMobile && r.Platform != PlatformPC {
		return ErrInvalidPlatform
	}
	if r.Currency.Code == "" || r.Currency.Name == "" {
		return ErrInvalidCurrency
	}
	return nil
}

func (r *SendBetRequest) Validate() error {
	if r.Token == "" {
		return ErrInvalidToken
	}
	if r.Player.Id <= 0 || r.Player.NickName == "" {
		return ErrInvalidPlayerParams
	}
	if r.GameID <= 0 {
		return ErrInvalidGameID
	}
	if r.Platform != PlatformMobile && r.Platform != PlatformPC {
		return ErrInvalidPlatform
	}
	if r.Currency.Code == "" || r.Currency.Name == "" {
		return ErrInvalidCurrency
	}
	if r.Amount <= 0 {
		return ErrInvalidAmount
	}
	if r.TransactionID <= 0 {
		return ErrInvalidTransactionID
	}
	return nil
}
