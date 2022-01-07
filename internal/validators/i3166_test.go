package validators

import (
	"encoding/json"
	"testing"
)

type Location struct {
	Country I3166 `json:"country" validate:"required"`
}

func TestI3166(t *testing.T) {
	validP := []byte(`{"country": "PRY"}`)
	invalidP := []byte(`{"country": ["PRY"]}`)

	testCases := []struct {
		desc          string
		value         []byte
		isSuccessCase bool
	}{
		{
			desc:          "Should be unmarshall with PRY code",
			value:         validP,
			isSuccessCase: true,
		},
		{
			desc:          "Should be unmarshall returns AppError",
			value:         invalidP,
			isSuccessCase: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var a Location
			err := json.Unmarshal(tC.value, &a)
			if err == nil && !tC.isSuccessCase {
				t.Error("invalid payload parse correctly")
			}

			if tC.isSuccessCase {
				if err != nil {
					t.Error("valid payload cannot parse correctly")
				}

				if a.Country.Code != "PRY" || a.Country.Name != "Paraguay" {
					t.Errorf("expected %v, got %v", I3166{Code: "PRY", Name: "Paraguay"}, a.Country)
				}
			}
		})
	}
}
