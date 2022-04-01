package transfer

import (
	"context"
	"errors"

	"github.com/dapr/go-sdk/actor"
	"github.com/go-logr/logr"
	"github.com/google/uuid"

	"github.com/pkedy/megabank/pkg/api"
)

// This file contains business logic.

type TransferActor struct {
	actor.ServerImplBase
	log        logr.Logger
	activities api.Activities
}

func Factory(log logr.Logger, activities api.Activities) actor.Factory {
	return func() actor.Server {
		return NewTransferActor(log, activities)
	}
}

func (t *TransferActor) Type() string {
	return "megabank.v1.Transfer"
}

func NewTransferActor(log logr.Logger, activities api.Activities) *TransferActor {
	return &TransferActor{
		log:        log,
		activities: activities,
	}
}

func (t *TransferActor) TransferFunds(ctx context.Context,
	request *api.TransferRequest) (*api.TransactionResult, error) {
	if request.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	info := []interface{}{
		"txid", t.ID(),
		"amount", request.Amount,
		"from", request.From,
		"to", request.To}
	t.log.Info("TransferFunds", info...)

	// Try to withdraw funds from the source account.
	if err := t.activities.Withdraw(ctx, request.Amount, request.From); err != nil {
		t.log.Error(err, "could not withdraw from source account", info...)
		return nil, err // err is properly wrapped
	}

	// Attempt to deposit into the destination account.
	if err := t.activities.Deposit(ctx, request.Amount, request.To); err != nil {
		t.log.Error(err, "could not deposit into destination account", info...)

		// Compensate for failure by returning the funds to the source account.
		if err = t.activities.Deposit(ctx, request.Amount, request.From); err != nil {
			t.log.Error(err, "could not return amount into source account", info...)
		}
	}

	t.log.Info("Transfer succeeded", info...)
	return &api.TransactionResult{
		TransactionID: uuid.New(),
		ReferenceID:   t.ID(),
	}, nil
}
