package currencyrates

import (
	"errors"

	"github.com/sackslab/bender/internal/currencylayer"
)

var (
	ErrUnsupportedCurrency = errors.New("unssuported currency")
)

func GetConversion(rates *currencylayer.RatesResponse, amount float64, from, to string) (float64, error) {
	fromRate, err := getRate(rates, from)
	if err != nil {
		return 0.0, err
	}

	toRate, err := getRate(rates, to)
	if err != nil {
		return 0.0, err
	}

	fromInUSD := amount / fromRate
	return fromInUSD * toRate, nil
}

func getRate(rates *currencylayer.RatesResponse, currency string) (float64, error) {
	var rate float64 = 1.0
	if currency != "USD" {
		var found bool
		rate, found = rates.Quotes["USD"+currency]
		if !found {
			return 0.0, ErrUnsupportedCurrency
		}
	}

	return rate, nil
}
