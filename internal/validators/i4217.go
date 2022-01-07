package validators

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/sackslab/bender/internal/i4217"
	"github.com/sackslab/bender/internal/middlewares/apperror"
)

type I4217 i4217.ISO4217

func (i *I4217) UnmarshalJSON(data []byte) error {
	var val interface{}
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	items := reflect.ValueOf(val)
	switch items.Kind() {
	case reflect.String:
		i4217, found := i4217.ByName(items.String())
		if !found {
			return &apperror.AppError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("invalid currency code or type %s, expected ISO4217 values", items.String()),
			}
		}
		*i = I4217(i4217)

	default:
		return &json.UnmarshalTypeError{Value: "string", Type: items.Type()}
	}

	return nil
}
