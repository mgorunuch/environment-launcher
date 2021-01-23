package console

import (
	"os/exec"

	"github.com/mgorunuch/environment-launcher/appcontext"
	"github.com/mgorunuch/environment-launcher/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrNotValidShellCommand = errors.New("not valid shell command")
)

func RunCommandsAsync(ctx appcontext.Context, commands []config.Command) {
	for _, c := range commands {
		go RunShellCommand(ctx, c)
	}
}

func RunShellCommand(ctx appcontext.Context, command config.Command) error {
	ctx.Logger.Info("stating shell command",
		zap.String("name", command.Name), zap.String("shell", command.Shell.Full()))

	if len(command.Shell) < 1 {
		err := errors.WithStack(ErrNotValidShellCommand)
		ctx.Logger.Error("shell command start", zap.Error(err), zap.Any("command", command))

		return err
	}

	base, argc := command.Shell.SplitBaseArgc()

	cmd := exec.CommandContext(ctx.Ctx, base, argc...)
	if err := cmd.Run(); err != nil {
		ctx.Logger.Error("failed to run command", zap.Error(err))
		return err
	}

	ctx.Logger.Info("program successfully executed", zap.String("command", command.Shell.Full()))

	return nil
}
