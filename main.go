package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mgorunuch/environment-launcher/app"

	"github.com/mgorunuch/environment-launcher/appcontext"

	"github.com/mgorunuch/environment-launcher/config"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger, _ := zap.NewDevelopment()

	var wg sync.WaitGroup

	startApp(logger, ctx, &wg)

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		wg.Wait()
		termChan <- syscall.SIGTERM
	}()

	<-termChan
	cancel()
}

func startApp(logger *zap.Logger, ctx context.Context, wg *sync.WaitGroup) {
	cnf, err := config.InitConfig(logger, "")
	if err != nil {
		logger.Fatal("config", zap.Error(err))
	}

	appCtx := appcontext.Context{
		Ctx:    ctx,
		Logger: logger,
		Config: cnf,
	}

	wg.Add(1)

	err = app.Start(appCtx, wg)
	if err != nil {
		logger.Fatal("app start", zap.Error(err))
	}
}
