package account

import (
	"context"
	"errors"

	"github.com/dapr/go-sdk/actor"
	"github.com/go-logr/logr"
	"github.com/pkedy/megabank/pkg/api"
)

// This file contains business logic.

const (
	KeyBalance = "balance"
)

var (
	ErrInsufficientFunds = errors.New("insufficient_funds")
)

type AccountActor struct {
	actor.ServerImplBase
	log logr.Logger
}

func Factory(log logr.Logger) actor.Factory {
	return func() actor.Server {
		return NewTransferActor(log)
	}
}

func (t *AccountActor) Type() string {
	return "megabank.v1.Account"
}

func NewTransferActor(log logr.Logger) *AccountActor {
	return &AccountActor{
		log: log,
	}
}

func (a *AccountActor) Balance(ctx context.Context) (api.Currency, error) {
	a.log.Info("Balance", "id", a.ID())
	var balance api.Currency
	if err := a.getOptional(KeyBalance, &balance); err != nil {
		a.log.Error(err, "could not get balance")
		return 0, err
	}
	a.log.Info("Balance",
		"account", a.ID(),
		"amount", balance)

	return balance, nil
}

func (a *AccountActor) Withdraw(ctx context.Context, amount api.Currency) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	var balance api.Currency
	sm := a.GetStateManager()
	if err := a.getOptional(KeyBalance, &balance); err != nil {
		a.log.Error(err, "could not get balance")
		return err
	}
	if amount > balance {
		a.log.Error(ErrInsufficientFunds, "not enough money",
			"amount", amount,
			"balance", balance)
		return ErrInsufficientFunds
	}

	newBalance := balance - amount
	sm.Set(KeyBalance, &newBalance)
	a.log.Info("Withdraw",
		"account", a.ID(),
		"amount", amount,
		"oldBalance", balance,
		"newBalance", newBalance)

	return nil // sm.Save() <-- this is required for state to save in the state store.
}

func (a *AccountActor) Deposit(ctx context.Context, amount api.Currency) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	var balance api.Currency
	sm := a.GetStateManager()
	if err := a.getOptional(KeyBalance, &balance); err != nil {
		a.log.Error(err, "could not get balance")
		return err
	}

	newBalance := balance + amount
	sm.Set(KeyBalance, &newBalance)
	a.log.Info("Deposit",
		"account", a.ID(),
		"amount", amount,
		"oldBalance", balance,
		"newBalance", newBalance)

	return nil // sm.Save() <-- this is required for state to save in the state store.
}

func (a *AccountActor) getOptional(key string, dst interface{}) error {
	sm := a.GetStateManager()
	ok, err := sm.Contains(key)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	return sm.Get(key, dst)
}
