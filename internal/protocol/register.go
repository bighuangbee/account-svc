package protocol

import (
	pb "github.com/bighuangbee/account-svc/api/account/v1"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PbServer struct {
	Account pb.AccountServer
}

func (s *PbServer) RegisterHTTP(srv *http.Server) {
	pb.RegisterAccountHTTPServer(srv, s.Account)
}

func (s *PbServer) RegisterRPC(srv *grpc.Server) {
	pb.RegisterAccountServer(srv, s.Account)
}
