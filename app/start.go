package app

import (
	"context"
	"errors"
	"fmt"
	appComponent "github.com/aitu-leetcode-site/core/app/component"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
)

func (a *App) addStarter(i appComponent.Starter) {
	if isNil(i) {
		panic("nil starter")
	}
	a.starters = append(a.starters, i)
}

func (a *App) addCloser(i appComponent.Closer) {
	if isNil(i) {
		panic("nil closer")
	}
	a.closers = append(a.closers, i)
}

func (a *App) addComponent(component any) {
	if isNil(component) {
		panic("nil component")
	}
	switch cmpType := component.(type) {
	case appComponent.StarterCloser:
		a.addStarter(cmpType)
		a.closers = append(a.closers, cmpType)
	case appComponent.Starter:
		a.addStarter(cmpType)
	case appComponent.Closer:
		a.addCloser(cmpType)
	}
}

func (a *App) Start(ctx context.Context) error {
	if len(a.starters) == 0 {
		return errors.New("no starters")
	}
	var err error
	a.ctx = ctx
	a.gracefulShutdown()
	wg := &sync.WaitGroup{}
	for _, s := range a.starters {
		wg.Add(1)
		go func(s appComponent.Starter) {
			defer wg.Done()
			startErr := s.Start(ctx)
			if startErr != nil {
				err = startErr
			}
		}(s)
	}
	wg.Wait()
	if err != nil {
		a.logger.Errorf(ctx, "Failed to start all starters: %v", err)
		a.chans.needCloseNotifyCh <- struct{}{}
		a.waitClose()
		return err
	}
	fmt.Print("app")
	a.runStatus.setRunning()
	a.logger.Infof(ctx, "App started.")
	a.waitClose()
	close(a.chans.closeAppCh)
	a.logger.Infof(ctx, "App stopped.")
	return err
}

func (a *App) close() error {
	defer a.setClosed()
	for _, c := range a.closers {
		err := c.Close(a.ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) gracefulShutdown() {
	go func() {
		select {
		case <-a.chans.needCloseNotifyCh:
			closeErr := a.close()
			if closeErr != nil {
				a.logger.Infof(a.ctx, "CloseErr: %v", closeErr)
			}
			return
		}
	}()
	sc := make(chan os.Signal)
	go func() {
		signal.Notify(sc,
			syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT,
		)
		<-sc
		a.logger.Infof(a.ctx, "recieved SIGTERM, shutting down")
		defer func() {
			if r := recover(); r != nil {
				a.logger.Errorf(a.ctx, "panic: %v", r)
			}
		}()
		a.chans.needCloseNotifyCh <- struct{}{}
		close(a.chans.needCloseNotifyCh)
		close(sc)
	}()
}

func isNil(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}

func (a *App) waitClose() {
	for {
		select {
		case <-a.chans.closeAppCh:
			return
		}
	}
}

func (a *App) setClosed() {
	a.runStatus.sedClosed()
	a.chans.closeAppCh <- struct{}{}
}
