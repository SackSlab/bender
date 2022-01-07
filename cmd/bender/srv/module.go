package srv

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sackslab/bender/cmd/bender/config"
	"github.com/sackslab/bender/internal/middlewares/apperror"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	return r
}

func RegistModule() fx.Option {
	return fx.Options(
		fx.Provide(New),
		fx.Invoke(runServer),
	)
}

func runServer(lc fx.Lifecycle, srv *gin.Engine, logger *zap.Logger, conf *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			port := fmt.Sprintf(":%d", conf.AppPort)
			logger.Info(fmt.Sprintf("[Starting] - serve at %s", port))

			srv.Use(apperror.JSONAppErrorReporter(logger))
			go srv.Run(port)
			return nil
		},
	})
}
