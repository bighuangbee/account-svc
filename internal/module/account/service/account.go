package service

import (
	"context"
	pb "github.com/bighuangbee/account-svc/api/account/v1"
	"github.com/bighuangbee/account-svc/internal/conf"
	"github.com/bighuangbee/account-svc/internal/domain"
	"github.com/go-kratos/kratos/v2/log"
)

type AccountService struct {
	repo             domain.IAccountRepo
	logger           *log.Helper
	bc               *conf.Bootstrap
}

func NewAccountService(repo domain.IAccountRepo, logger log.Logger, bc *conf.Bootstrap)(*AccountService){
	return &AccountService{
		repo:   nil,
		logger: nil,
		bc:     nil,
	}
}


func (this *AccountService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
