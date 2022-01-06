package main

import (
	"github.com/sackslab/bender/internal/logger"
	"go.uber.org/fx"
)

// TODO: setup db conn
// TODO: setup gin server
func main() {
	app := fx.New(
		logger.RegistModule(),
	)

	app.Run()
}
