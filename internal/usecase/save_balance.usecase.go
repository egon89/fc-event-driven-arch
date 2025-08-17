package usecase

import (
	"context"
	"fmt"

	"github.com/egon89/fc-event-driven-arch/internal/entity"
	"github.com/egon89/fc-event-driven-arch/internal/service"
)

type SaveBalanceInputDto struct {
	AccountIdFrom      string  `json:"accountIdFrom"`
	BalanceAccountFrom float64 `json:"balanceAccountFrom"`
	AccountIdTo        string  `json:"accountIdTo"`
	BalanceAccountTo   float64 `json:"balanceAccountTo"`
}

type SaveBalanceUseCase struct {
	balanceService service.BalanceService
}

func NewSaveBalanceUseCase(balanceService service.BalanceService) *SaveBalanceUseCase {
	return &SaveBalanceUseCase{balanceService}
}

func (uc *SaveBalanceUseCase) Execute(ctx context.Context, input SaveBalanceInputDto) error {
	err := uc.upsertBalance(ctx, input.AccountIdFrom, input.BalanceAccountFrom)
	if err != nil {
		return fmt.Errorf("error saving balance for account %s: %v", input.AccountIdFrom, err)
	}

	err = uc.upsertBalance(ctx, input.AccountIdTo, input.BalanceAccountTo)
	if err != nil {
		return fmt.Errorf("error saving balance for account %s: %v", input.AccountIdTo, err)
	}

	return nil
}

func (uc *SaveBalanceUseCase) upsertBalance(ctx context.Context, accountId string, balanceAmount float64) error {
	balance, err := uc.balanceService.FindOneByAccountId(ctx, accountId)
	if err != nil {
		return fmt.Errorf("error finding balance for account %s: %v", accountId, err)
	}

	if balance == nil {
		return uc.balanceService.Save(ctx, &entity.Balance{
			AccountId: accountId,
			Balance:   balanceAmount,
		})
	}

	balance.Balance += balanceAmount
	return uc.balanceService.Update(ctx, balance)
}
