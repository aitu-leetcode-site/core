package app

import (
	"context"
	"github.com/aitu-leetcode-site/core/app/component"
	"github.com/aitu-leetcode-site/core/config"
	"github.com/aitu-leetcode-site/core/grpc"
	"github.com/aitu-leetcode-site/core/http"
	httpconfig "github.com/aitu-leetcode-site/core/http/config"
	"github.com/aitu-leetcode-site/core/log"
	"github.com/aitu-leetcode-site/core/utils"
	"time"
)

type App struct {
	noCopy     noCopy
	ctx        context.Context
	cfg        *config.Config
	starters   []component.Starter
	closers    []component.Closer
	httpServer *http.Server
	grpcServer *grpc.GRPSServer
	runStatus  runStatus
	chans      struct {
		needCloseNotifyCh chan struct{}
		closeAppCh        chan struct{}
	}
	logger log.Logger
}

func NewApp(appName string, cfg *config.Config) *App {
	app := &App{
		cfg: cfg,
		chans: struct {
			needCloseNotifyCh chan struct{}
			closeAppCh        chan struct{}
		}{
			needCloseNotifyCh: make(chan struct{}),
			closeAppCh:        make(chan struct{}),
		},
		logger: log.GetLogger(),
	}
	httpS := http.NewServer(appName, &httpconfig.Config{
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Port:         utils.Port(8080),
	})
	app.httpServer = httpS
	app.addComponent(httpS)
	grpcS := grpc.NewGRPCServer()
	app.grpcServer = grpcS
	app.addComponent(grpcS)
	return app
}
