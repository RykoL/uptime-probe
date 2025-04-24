package probe_test

import (
	"github.com/RykoL/uptime-probe/internal/monitor/probe"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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
	targetUrl := server.URL + "/success"
	defer server.Close()

	httpProbe := probe.NewHttpProbe(targetUrl)

	result, err := httpProbe.Execute()

	assert.Nil(t, err)
	assert.Equal(t, result.Succeeded, probe.ExecutionSucceeded)
}

func TestReturnsFailureOnNonSuccessfulRequest(t *testing.T) {

	server := server()
	targetUrl := server.URL + "/failure"
	defer server.Close()

	httpProbe := probe.NewHttpProbe(targetUrl)

	result, err := httpProbe.Execute()

	assert.Nil(t, err)
	assert.Equal(t, result.Succeeded, probe.ExecutionFailed)
}
