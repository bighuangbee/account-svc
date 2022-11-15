// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/bighuangbee/account-svc/internal/conf"
	"github.com/bighuangbee/account-svc/internal/data"
	"github.com/bighuangbee/account-svc/internal/module/account"
	"github.com/bighuangbee/account-svc/internal/protocol"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/bighuangbee/gokit/log/kitZap"
)

func autoWireApp(*conf.Bootstrap, log.Logger, *kitZap.ZapLogger, *i18n.Bundle) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, protocol.ProviderSet, account.ProviderSet, newApp))
}
