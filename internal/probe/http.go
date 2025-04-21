package probe

import (
	"net/http"
	"net/url"
)

type HttpProbe struct {
	url *url.URL
}

func NewHttpProbe(url *url.URL) HttpProbe {
	return HttpProbe{url}
}

func (p *HttpProbe) Execute() (*Result, error) {
	resp, err := http.Get(p.url.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return &Result{Status: StatusSucceed}, nil
	}

	return &Result{Status: StatusFailed}, nil
}
