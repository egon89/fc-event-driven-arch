package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/egon89/fc-event-driven-arch/internal/entity"
	"github.com/egon89/fc-event-driven-arch/internal/gateway"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type balanceModel struct {
	Id        string    `db:"id"`
	AccountId string    `db:"account_id"`
	Balance   float64   `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type balanceRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) gateway.BalanceGateway {
	return &balanceRepository{db}
}

func (b *balanceRepository) FindOneByAccountId(ctx context.Context, accountId string) (*entity.Balance, error) {
	query := `
		SELECT id, account_id, balance, created_at, updated_at
		FROM balance
		WHERE account_id = $1
		LIMIT 1
	`
	var balance balanceModel
	err := b.db.GetContext(ctx, &balance, query, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &entity.Balance{
		Id:        balance.Id,
		AccountId: balance.AccountId,
		Balance:   balance.Balance,
		CreatedAt: balance.CreatedAt,
		UpdatedAt: balance.UpdatedAt,
	}, nil
}

func (b *balanceRepository) Save(ctx context.Context, balance *entity.Balance) error {
	if balance.Id == "" {
		balance.Id = uuid.NewString()
	}

	query := `
        INSERT INTO balance (id, account_id, balance, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	return b.db.QueryRowxContext(ctx, query,
		balance.Id,
		balance.AccountId,
		balance.Balance,
		balance.CreatedAt,
		balance.UpdatedAt,
	).Scan(&balance.Id)
}

func (b *balanceRepository) Update(ctx context.Context, balance *entity.Balance) error {
	query := `
		UPDATE balance
		SET balance = $1, updated_at = $2
		WHERE account_id = $3
	`
	balance.UpdatedAt = time.Now()

	_, err := b.db.ExecContext(ctx, query,
		balance.Balance,
		balance.UpdatedAt,
		balance.AccountId,
	)

	if err != nil {
		return err
	}

	return nil
}
