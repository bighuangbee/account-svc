package domain

import (
	"context"
)

type Account struct {
	Hello string
}

type IAccountRepo interface {
	Login(context.Context, *Account) (*Account, error)
}
