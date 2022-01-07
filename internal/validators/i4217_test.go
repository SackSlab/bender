package validators

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/sackslab/bender/internal/middlewares/apperror"
)

type Amount struct {
	Currency I4217 `json:"currency" validate:"required"`
}

func TestI4217(t *testing.T) {
	validP := []byte(`{"currency": "PYG"}`)
	invalidPType := []byte(`{"currency": ["PYG"]}`)
	invalidPCode := []byte(`{"currency": "HHSS"}`)

	testCases := []struct {
		desc          string
		value         []byte
		isSuccessCase bool
	}{
		{
			desc:          "Should be unmarshall returns AppError for invalid type",
			value:         invalidPType,
			isSuccessCase: false,
		},
		{
			desc:          "Should be unmarshall returns AppError for invalid code",
			value:         invalidPCode,
			isSuccessCase: false,
		},
		{
			desc:          "Should be unmarshall with PYG code",
			value:         validP,
			isSuccessCase: true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var a Amount
			err := json.Unmarshal(tC.value, &a)
			if !tC.isSuccessCase {
				if err == nil {
					t.Error("invalid payload parse correctly")
				}

				if aerr, ok := err.(*apperror.AppError); ok && aerr.Code != http.StatusBadRequest {
					t.Errorf("expected 400 - bad request code")
				}
			}

			if tC.isSuccessCase {
				if err != nil {
					t.Error("valid payload cannot parse correctly")
				}

				if a.Currency.Code != "600" || a.Currency.Name != "PYG" || a.Currency.Units != 0 {
					t.Errorf("expected %v, got %v", I4217{Code: "600", Name: "PYG", Units: 0}, a.Currency)
				}
			}
		})
	}
}
