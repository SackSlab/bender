package currencylayer

import (
	"testing"
)

var opts = Options{
	HostURL: "http://api.currencylayer.com",
	// replace with your api key
	ApiKey: "0417d6f26c8212b21109a33120b7a1e2",
}

func TestCall(t *testing.T) {
	c := NewClient(opts)

	resp, err := c.Latest()
	if err != nil {
		t.Errorf("unexpected error when fetching latest currency rates, %s", err)
	}

	_ = resp
}
