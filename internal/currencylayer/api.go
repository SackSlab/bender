package currencylayer

import (
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

type Options struct {
	HostURL string
	ApiKey  string
}

type Client struct {
	client *http.Client
	opts   Options
}

func NewClient(opts Options) *Client {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	return &Client{
		client: &http.Client{Jar: cookieJar},
		opts:   opts,
	}
}

func (c *Client) Latest() (*RatesResponse, error) {
	url, err := c.prepareURL("live")
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cResp RatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return nil, err
	}

	return &cResp, nil
}

func (c *Client) prepareURL(p string) (*url.URL, error) {
	base, err := url.Parse(c.opts.HostURL)
	if err != nil {
		return nil, err
	}

	path, _ := url.Parse(p)
	resolved := base.ResolveReference(path)
	qry := resolved.Query()
	qry.Add("access_key", c.opts.ApiKey)
	resolved.RawQuery = qry.Encode()

	return resolved, nil
}
