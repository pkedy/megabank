package transfer

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/pkedy/megabank/pkg/api"
)

// This file contains business logic.

type ActivitiesImpl struct {
	daprClient dapr.Client
}

var _ = (api.Activities)((*ActivitiesImpl)(nil))

func NewActivities(daprClient dapr.Client) *ActivitiesImpl {
	return &ActivitiesImpl{
		daprClient: daprClient,
	}
}

func (a *ActivitiesImpl) Withdraw(ctx context.Context, amount api.Currency, from string) error {
	acccount := api.NewAccountActor(from)
	a.daprClient.ImplActorClientStub(acccount)

	// Resiliency is configured elsewhere in a CRD
	return acccount.Withdraw(ctx, amount)
}

func (a *ActivitiesImpl) Deposit(ctx context.Context, amount api.Currency, to string) error {
	acccount := api.NewAccountActor(to)
	a.daprClient.ImplActorClientStub(acccount)

	// Resiliency is configured elsewhere in a CRD
	return acccount.Deposit(ctx, amount)
}
