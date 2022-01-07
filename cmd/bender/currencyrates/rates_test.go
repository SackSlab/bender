package currencyrates

import (
	"errors"
	"testing"

	"github.com/sackslab/bender/internal/currencylayer"
)

func Test(t *testing.T) {
	rates := &currencylayer.RatesResponse{
		Success: true,
		Source:  "USD",
		Quotes: map[string]float64{
			"USDARS": 103.230994,
			"USDPYG": 6944.249037,
		},
	}

	testCases := []struct {
		desc           string
		expectedAmount float64
		amount         float64
		from           string
		to             string
		isSuccessCase  bool
	}{
		{
			desc:           "Should be convert correctly",
			expectedAmount: 178.388177238408,
			amount:         12000.00,
			from:           "PYG",
			to:             "ARS",
			isSuccessCase:  true,
		},
		{
			desc:           "Should be convert correctly with USD",
			expectedAmount: 13888498.074,
			amount:         2000.00,
			from:           "USD",
			to:             "PYG",
			isSuccessCase:  true,
		},
		{
			desc:           "Should be return unssuported error",
			expectedAmount: 13888498.074,
			amount:         2000.00,
			from:           "SWC",
			to:             "PYG",
			isSuccessCase:  false,
		},
		{
			desc:           "Should be return unssuported error for 'to' param",
			expectedAmount: 13888498.074,
			amount:         2000.00,
			to:             "SWC",
			from:           "PYG",
			isSuccessCase:  false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			conv, err := GetConversion(rates, tC.amount, tC.from, tC.to)
			if err != nil && errors.Is(err, ErrUnsupportedCurrency) && tC.isSuccessCase {
				t.Error("unexpected error")
			}

			if tC.isSuccessCase && conv != tC.expectedAmount {
				t.Errorf("expected value its %.2f, got %.2f", tC.expectedAmount, conv)
			}
		})
	}
}
