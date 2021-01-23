package appcontext

import (
	"context"

	"github.com/mgorunuch/environment-launcher/config"

	"go.uber.org/zap"
)

type Context struct {
	Ctx    context.Context
	Config *config.Config
	Logger *zap.Logger
}
