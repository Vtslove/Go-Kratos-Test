package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-blog/gen/api/go/common/conf"
	"kratos-blog/pkg/bootstrap"
	"kratos-blog/pkg/service"
)

// go build -ldflags "-X main.Service.Version=x.y.z"
var (
	Service = bootstrap.NewServiceInfo(
		service.ContentService,
		"1.0.0",
		"",
	)

	Flags = bootstrap.NewCommandFlags()
)

func init() {
	Flags.Init()
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rr),
	)
}

func loadConfig() (*conf.Bootstrap, *conf.Registry) {
	c := bootstrap.NewConfigProvider(Flags.ConfigType, Flags.ConfigHost, Flags.Conf, Service.Name)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	return &bc, &rc
}

func main() {
	flag.Parse()

	bc, rc := loadConfig()
	if bc == nil || rc == nil {
		panic("load config failed")
	}

	logger := bootstrap.NewLoggerProvider(bootstrap.LoggerTypeStd, bc.Logger, &Service)

	err := bootstrap.NewTracerProvider(bc.Trace, &Service)
	if err != nil {
		panic(err)
	}

	app, cleanupApp, err := initApp(bc.Server, rc, bc.Data, bc.Auth, logger)
	if err != nil {
		panic(err)
	}
	defer cleanupApp()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
