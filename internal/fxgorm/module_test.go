package fxgorm

import (
	"fmt"
	"reflect"
	"testing"

	pg "gorm.io/driver/postgres"
)

const (
	dsn = "anotatedddd"
	// TODO: use snapshots for modules signatures
	moduleSignature = `fx.Options(fx.Provide(fx.Annotate(github.com/sackslab/bender/internal/fxgorm.dsnFunc(), fx.ResultTags(["name:\"gorm_dsn\""])), fx.Provide(fx.Annotate(github.com/sackslab/bender/internal/fxgorm.NewPGDialector(), fx.ParamTags(["name:\"gorm_dsn\""])), fx.Provide(gorm.io/gorm.Open()))`
)

func TestDSNAnotation(t *testing.T) {
	annotated := DSNAnnotation(dsnFunc)

	rv := reflect.ValueOf(annotated)
	rTagVal := rv.FieldByName("ResultTags")
	rTargetVal, ok := rv.FieldByName("Target").Interface().(func() string)
	if !ok {
		t.Error("dsnAnnotation its not func() string")
	}

	if val := rTargetVal(); val != dsn {
		t.Errorf("expected value its '%s', got '%s'", dsn, val)
	}

	if rTagStr := fmt.Sprintf("%v", rTagVal); rTagStr != `[name:"gorm_dsn"]` {
		t.Errorf("invalid result tag, got %s", rTagStr)
	}
}

func TestNewPGDialector(t *testing.T) {
	dialector := NewPGDialector(dsnFunc())
	if (&pg.Dialector{}).Name() != dialector.Name() {
		t.Error("invalid dialector")
	}
}

func TestNewModule(t *testing.T) {
	mod := RegistModule(dsnFunc)
	if mod.String() != moduleSignature {
		t.Errorf("module signature its invalid")
	}
}

func dsnFunc() string {
	return dsn
}
