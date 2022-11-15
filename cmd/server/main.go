package main

import (
	"flag"
	"fmt"
	"github.com/bighuangbee/gokit/kitGoi18n"
	"github.com/bighuangbee/gokit/log/kitZap"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	clientv3 "go.etcd.io/etcd/client/v3"
	"path"
	"time"

	"github.com/bighuangbee/account-svc/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.uber.org/zap/zapcore"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	_ "net/http/pprof"
	grpcDial "google.golang.org/grpc"
)

var flagConf string

func newApp(bc *conf.Bootstrap, logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	var registrar registry.Registrar
	if bc.Discovery.OnOff {
		client, err := clientv3.New(clientv3.Config{
			Endpoints:   bc.MicroService.Etcd.Addr,
			DialTimeout: time.Second, DialOptions: []grpcDial.DialOption{grpcDial.WithBlock()},
		})
		if err != nil {
			panic(err)
		}
		registrar = etcd.New(client)
	}

	return kratos.New(
		kratos.Name(bc.Name),
		kratos.Version(bc.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(hs, gs),
		kratos.Registrar(registrar),
	)
}

func main() {
	flag.StringVar(&flagConf, "conf", "../../config/config.dev.yaml", "config path, eg: -conf config.yaml")
	flag.Parse()

	var bc *conf.Bootstrap

	fmt.Println("----flagConf", flagConf)
	c := config.New(config.WithSource(file.NewSource(flagConf)))
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// zap 开启stack trace logger, 用于service层输出
	zapLog := kitZap.New(&kitZap.Options{
		Level: zapcore.DebugLevel,
		Skip:  2,
	})
	logger := log.With(zapLog, "tid", tracing.TraceID())

	langPath := path.Dir(flagConf) + "/i18n/"
	bundle := kitGoi18n.New(kitGoi18n.Options{
		Paths: []string{
			langPath + "example.en.toml",
			langPath + "example.zh.toml",
		},
	})
	//opLogCli := toolsMiddleware.NewOpLog(path.Dir(flagConf), bc.OpLog.OpLogGrpcAddr, bc.ServerItem.AppId, logger)
	app, cleanup, err := autoWireApp(bc, logger, zapLog, bundle)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
