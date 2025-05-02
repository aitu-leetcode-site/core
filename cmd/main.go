package main

import (
	"context"
	coreApp "github.com/aitu-leetcode-site/core/app"
	"github.com/aitu-leetcode-site/core/config"
	"github.com/aitu-leetcode-site/core/log"
)

type TestApp struct {
	*coreApp.App
}

func InitApp() *TestApp {
	return &TestApp{
		App: coreApp.NewApp("test", &config.Config{}),
	}
}

func main() {
	app := InitApp()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		log.Panic(ctx, err.Error())
	}
}
