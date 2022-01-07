package currencylayer

import "go.uber.org/fx"

func RegistModule(optionsProvider interface{}) fx.Option {
	return fx.Options(
		fx.Provide(
			optionsProvider,
			NewClient,
		),
	)
}
