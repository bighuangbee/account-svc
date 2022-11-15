package account

import (
	"github.com/bighuangbee/account-svc/internal/module/account/app"
	"github.com/bighuangbee/account-svc/internal/module/account/repo"
	"github.com/bighuangbee/account-svc/internal/module/account/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(app.NewAccountApp, service.NewAccountService, repo.NewAccountRepo)
