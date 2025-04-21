package probe_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"uptime-probe/internal/probe"
)

func server() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /success", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("GET /fail", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	})

	return httptest.NewServer(mux)
}

func TestReturnsSuccessfulHttpResult(t *testing.T) {

	server := server()
	targetUrl, _ := url.Parse(server.URL)
	defer server.Close()

	httpProbe := probe.NewHttpProbe(targetUrl.JoinPath("/success"))

	result, err := httpProbe.Execute()

	assert.Nil(t, err)
	assert.Equal(t, result.Status, probe.StatusSucceed)
}

func TestReturnsFailureOnNonSuccessfulRequest(t *testing.T) {

	server := server()
	targetUrl, _ := url.Parse(server.URL)
	defer server.Close()

	httpProbe := probe.NewHttpProbe(targetUrl.JoinPath("/failure"))

	result, err := httpProbe.Execute()

	assert.Nil(t, err)
	assert.Equal(t, result.Status, probe.StatusFailed)
}
