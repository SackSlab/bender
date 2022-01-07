package logger

import "testing"

const moduleSignature = "fx.Options(fx.Provide(go.uber.org/zap.NewProduction()), fx.WithLogger(github.com/sackslab/bender/internal/logger.useZapLogger()))"

func TestNewModule(t *testing.T) {
	mod := RegistModule()
	if mod.String() != moduleSignature {
		t.Errorf("module signature its invalid")
	}
}
