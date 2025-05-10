package probe

import (
	"encoding/json"
	"fmt"
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
		return &ProbeResult{Succeeded: ExecutionFailed}, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return &ProbeResult{Succeeded: ExecutionSucceeded}, nil
	}

	return &ProbeResult{Succeeded: ExecutionFailed}, nil
}

func (p *HttpProbe) AsJSON() (string, error) {
	bytes, err := json.Marshal(p)

	if err != nil {
		return "", fmt.Errorf("failed to marshal http probe %w", err)
	}

	return string(bytes), nil
}
