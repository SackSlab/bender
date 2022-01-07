package fxgorm

import (
	"go.uber.org/fx"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsnParamKey = `name:"gorm_dsn"`

func DSNAnnotation(dsnFunc interface{}) interface{} {
	return fx.Annotate(
		dsnFunc,
		fx.ResultTags(dsnParamKey),
	)
}

func NewPGDialector(dsn string) gorm.Dialector {
	return pg.New(pg.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
}

func RegistModule(dsnProviderFunc interface{}) fx.Option {
	return fx.Options(
		fx.Provide(DSNAnnotation(dsnProviderFunc)),
		fx.Provide(fx.Annotate(
			NewPGDialector,
			fx.ParamTags(dsnParamKey),
		)),
		fx.Provide(gorm.Open),
	)
}
