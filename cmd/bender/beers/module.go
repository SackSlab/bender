package beers

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegistModule() fx.Option {
	return fx.Options(
		fx.Provide(
			NewService,
			NewController,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, r *gin.Engine, ctrl *controller) {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						ctrl.ConfigureRouter(r)
						return nil
					},
				})
			},
		),
	)
}
