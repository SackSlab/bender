package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func useZapLogger(logger *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: logger}
}

func RegistModule() fx.Option {
	return fx.Options(
		fx.Provide(zap.NewProduction),
		fx.WithLogger(useZapLogger),
	)
}
