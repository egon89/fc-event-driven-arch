package entity

import "time"

type Balance struct {
	Id        string
	AccountId string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
