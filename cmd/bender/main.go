package main

import (
	"fmt"

	"github.com/sackslab/bender/cmd/bender/beers"
	"github.com/sackslab/bender/cmd/bender/config"
	"github.com/sackslab/bender/cmd/bender/srv"
	"github.com/sackslab/bender/internal/fxgorm"
	"github.com/sackslab/bender/internal/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		logger.RegistModule(),
		config.RegistModule(),
		fxgorm.RegistModule(
			func(c *config.Config) string {
				return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
					c.DBHost, c.DBUser, c.DBPass, c.DBName)
			},
		),
		srv.RegistModule(),
		beers.RegistModule(),
	)

	app.Run()
}
