package probe

import (
	"net/http"
)

type HttpProbe struct {
	Url string `json:"Url"`
}

func NewHttpProbe(url string) *HttpProbe {
	return &HttpProbe{url}
}

func (p *HttpProbe) Target() string {
	return p.Url
}

func (p *HttpProbe) Execute() (*ProbeResult, error) {
	resp, err := http.Get(p.Url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return &ProbeResult{Succeeded: ExecutionSucceeded}, nil
	}

	return &ProbeResult{Succeeded: ExecutionFailed}, nil
}
