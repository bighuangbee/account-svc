package protocol

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bighuangbee/account-svc/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(bc *conf.Bootstrap, logger log.Logger, services *PbServer) *grpc.Server {
	c := bc.Server

	srv := grpc.NewServer(
		grpc.Address(c.Grpc.Addr),
		grpc.Timeout(time.Duration(c.Grpc.Timeout)*time.Second),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			validate(),
		),
		grpc.Logger(logger),
	)

	services.RegisterRPC(srv)
	return srv
}

type Validate interface {
	Validate() error
}

func validate(opts ...recovery.Option) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if _, ok := req.(emptypb.Empty); ok {
				return handler(ctx, req)
			}
			if _, ok := req.(*emptypb.Empty); ok {
				return handler(ctx, req)
			}
			r := req.(Validate)
			if err := r.Validate(); err != nil {
				return nil, err
			}
			return handler(ctx, req)
		}
	}
}
