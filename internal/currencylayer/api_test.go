package currencylayer

import (
	"testing"
)

var opts = Options{
	HostURL: "http://api.currencylayer.com",
	// replace with your api key
	ApiKey: "API_KEY",
}

func TestCall(t *testing.T) {
	c := NewClient(opts)

	resp, err := c.Latest()
	if err != nil {
		t.Errorf("unexpected error when fetching latest currency rates, %s", err)
	}

	_ = resp
}
