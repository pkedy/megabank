package api

import (
	"context"
)

type AccountActor struct {
	id       string
	Balance  func(ctx context.Context) (Currency, error)
	Withdraw func(ctx context.Context, amount Currency) error
	Deposit  func(ctx context.Context, amount Currency) error
}

type AmountRequest struct {
	Amount Currency `json:"amount" type:"integer" format:"int64"`
}

func NewAccountActor(id string) *AccountActor {
	return &AccountActor{
		id: id,
	}
}

func (a *AccountActor) Type() string {
	return "megabank.v1.Account"
}

func (a *AccountActor) ID() string {
	return a.id
}
