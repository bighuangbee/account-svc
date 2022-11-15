package app

import (
	"context"
	pb "github.com/bighuangbee/account-svc/api/account/v1"
	"github.com/bighuangbee/account-svc/internal/module/account/service"
	"github.com/bighuangbee/gokit/log/kitZap"
)

type AccountApp struct {
	pb.UnimplementedAccountServer
	svc    *service.AccountService
	logger *kitZap.ZapLogger
}

func NewAccountApp(svc *service.AccountService, logger *kitZap.ZapLogger) pb.AccountServer {

	return &AccountApp{
		svc:    svc,
		logger: logger,
	}
}

func (s *AccountApp) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	return s.svc.Login(ctx, req)
}
