package api

import (
	"context"

	"github.com/google/uuid"
)

type TransferActor struct {
	id            string
	TransferFunds func(ctx context.Context, request *TransferRequest) (*TransactionResult, error)
}

type TransferRequest struct {
	Amount Currency `json:"amount"`
	From   string   `json:"from"`
	To     string   `json:"to"`
}

type TransactionResult struct {
	TransactionID uuid.UUID `json:"transactionId"`
	ReferenceID   string    `json:"referenceId"`
}

type Activities interface {
	Withdraw(ctx context.Context, amount Currency, from string) error
	Deposit(ctx context.Context, amount Currency, to string) error
}

func NewTransferActor(id string) *TransferActor {
	return &TransferActor{
		id: id,
	}
}

func (a *TransferActor) Type() string {
	return "megabank.v1.Transfer"
}

func (a *TransferActor) ID() string {
	return a.id
}
