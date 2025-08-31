package usecase

import (
	"context"
	"fmt"

	"github.com/egon89/fc-event-driven-arch/internal/service"
)

type FindBalanceByAccountIdOutputDto struct {
	Id        string
	AccountId string
	Balance   float64
}

type FindBalanceByAccountIdUseCase struct {
	balanceService service.BalanceService
}

func NewFindBalanceByAccountIdUseCase(balanceService service.BalanceService) *FindBalanceByAccountIdUseCase {
	return &FindBalanceByAccountIdUseCase{balanceService}
}

func (uc *FindBalanceByAccountIdUseCase) Execute(ctx context.Context, id string) (*FindBalanceByAccountIdOutputDto, error) {
	balance, err := uc.balanceService.FindOneByAccountId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find balance by account id: %w", err)
	}

	if balance == nil {
		return nil, fmt.Errorf("balance with id %v not found", id)
	}

	output := &FindBalanceByAccountIdOutputDto{
		Id:        balance.Id,
		AccountId: balance.AccountId,
		Balance:   balance.Balance,
	}

	return output, nil
}
