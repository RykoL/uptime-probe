package monitor_test

import (
	"github.com/RykoL/uptime-probe/internal/monitor"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
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

	httpProbe := monitor.NewHttpProbe(targetUrl.JoinPath("/success"))

	result, err := httpProbe.Execute()

	assert.Nil(t, err)
	assert.Equal(t, result.Succeeded, monitor.ExecutionSucceeded)
}

func TestReturnsFailureOnNonSuccessfulRequest(t *testing.T) {

	server := server()
	targetUrl, _ := url.Parse(server.URL)
	defer server.Close()

	httpProbe := monitor.NewHttpProbe(targetUrl.JoinPath("/failure"))

	result, err := httpProbe.Execute()

	assert.Nil(t, err)
	assert.Equal(t, result.Succeeded, monitor.ExecutionFailed)
}
