package repo

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/bighuangbee/account-svc/internal/conf"
	"github.com/bighuangbee/account-svc/internal/data"
	"github.com/bighuangbee/account-svc/internal/domain"
)


func NewAccountRepo(data *data.Data, logger log.Logger, bootstrap *conf.Bootstrap) domain.IAccountRepo {
	return &AccountRepo{
		data:   data,
		logger: log.NewHelper(logger),
		bc:     bootstrap,
	}
}

type AccountRepo struct {
	data   *data.Data
	logger *log.Helper
	bc     *conf.Bootstrap
}

func (this *AccountRepo) Login(context.Context, *domain.Account)  (*domain.Account, error) {
	return &domain.Account{}, nil
}
