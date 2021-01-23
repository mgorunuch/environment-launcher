package app

import (
	"sync"

	"github.com/mgorunuch/environment-launcher/appcontext"
	"github.com/mgorunuch/environment-launcher/console"
	"github.com/mgorunuch/environment-launcher/tray"
	"go.uber.org/zap"
)

func Start(ctx appcontext.Context, wg *sync.WaitGroup) error {
	trayApps := tray.Apps{}

	for _, app := range ctx.Config.Apps {
		trayApps[app.Name] = func() {
			ctx.Logger.Info("started app", zap.String("name", app.Name))
			console.RunCommandsAsync(ctx, app.Commands)
		}
	}

	err := tray.Start(ctx, wg, trayApps)
	if err != nil {
		return err
	}

	return nil
}
