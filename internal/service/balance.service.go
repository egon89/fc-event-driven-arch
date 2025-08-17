package service

import (
	"context"

	"github.com/egon89/fc-event-driven-arch/internal/entity"
	"github.com/egon89/fc-event-driven-arch/internal/gateway"
)

type balanceServiceImp struct {
	gateway gateway.BalanceGateway
}

// In Go, you return a pointer to a struct that implements an interface,
// but the function signature uses the interface type, not a pointer to the interface.
// Interfaces are reference types, so you never use *InterfaceType.
// The pointer to the struct is assigned to the interface, which holds a reference to it.
// Go interfaces are satisfied by any type (pointer or value) that implements the interface methods.
func NewBalanceService(gateway gateway.BalanceGateway) BalanceService {
	return &balanceServiceImp{gateway}
}

func (s *balanceServiceImp) FindOneByAccountId(ctx context.Context, accountId string) (*entity.Balance, error) {
	return s.gateway.FindOneByAccountId(ctx, accountId)
}

func (s *balanceServiceImp) Save(ctx context.Context, balance *entity.Balance) error {
	return s.gateway.Save(ctx, balance)
}

func (s *balanceServiceImp) Update(ctx context.Context, balance *entity.Balance) error {
	return s.gateway.Update(ctx, balance)
}
