package service

import (
	"context"

	"github.com/egon89/fc-event-driven-arch/internal/entity"
)

type BalanceService interface {
	FindOneByAccountId(ctx context.Context, accountId string) (*entity.Balance, error)
	Save(ctx context.Context, balance *entity.Balance) error
	Update(ctx context.Context, balance *entity.Balance) error
}
