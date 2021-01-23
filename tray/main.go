package tray

import (
	"context"
	"sync"

	"github.com/mgorunuch/environment-launcher/config"

	"github.com/mgorunuch/environment-launcher/console"

	"github.com/mgorunuch/environment-launcher/appcontext"

	"go.uber.org/zap"

	"github.com/mgorunuch/environment-launcher/icons"

	"github.com/getlantern/systray"
)

type Apps map[string]func()

func Start(ctx appcontext.Context, wg *sync.WaitGroup, apps Apps) (err error) {
	defer func() {
		if err != nil {
			wg.Done()
		}
	}()

	localCtx, cancel := context.WithCancel(ctx.Ctx)

	shuttleIcon, err := icons.GetShuttleIcon()
	if err != nil {
		ctx.Logger.Error("get icon", zap.Error(err))
		return err
	}

	exitIcon, err := icons.GetExitIcon()
	if err != nil {
		ctx.Logger.Error("get icon", zap.Error(err))
		return err
	}

	go systray.Run(func() {
		systray.SetIcon(shuttleIcon)

		for appName, app := range apps {
			item := systray.AddMenuItem(appName, "")
			app := app
			appName := appName
			go func() {
				for {
					<-item.ClickedCh
					ctx.Logger.Info("run app", zap.String("app", appName))
					app()
				}
			}()
		}

		mEdit := systray.AddMenuItem("Edit config", "Edit program config")
		go func() {
			for {
				<-mEdit.ClickedCh
				go console.RunShellCommand(ctx, config.Command{
					Name:  "file editor",
					Shell: config.ShellCommand{"xdg-open", ctx.Config.ConfigPath},
				})
			}
		}()

		mQuit := systray.AddMenuItem("Exit", "Exit the whole app")
		mQuit.SetIcon(exitIcon)
		go func() {
			<-mQuit.ClickedCh
			cancel()
		}()

		go func() {
			<-localCtx.Done()
			systray.Quit()
		}()

		ctx.Logger.Info("systray started")
	}, func() {
		wg.Done()
		ctx.Logger.Info("systray exited")
	})

	return nil
}
