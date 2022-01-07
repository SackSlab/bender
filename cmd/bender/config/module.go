package config

import (
	"github.com/Netflix/go-env"
	"go.uber.org/fx"
)

type Config struct {
	DBName string `env:"POSTGRES_DB,required=true"`
	DBPass string `env:"POSTGRES_PASSWORD,required=true"`
	DBUser string `env:"POSTGRES_USER,required=true"`
	DBHost string `env:"POSTGRES_HOST,required=true"`

	AppPort int `env:"APP_PORT,required=true"`

	CurrencyLayerApiURL    string `env:"CL_API_URL,required=true"`
	CurrencyLayerAccessKey string `env:"CL_ACCESS_KEY,required=true"`
}

func New() (*Config, error) {
	var conf Config
	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func RegistModule() fx.Option {
	return fx.Provide(New)
}
