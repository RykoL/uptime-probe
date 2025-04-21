package monitor

import (
	"github.com/RykoL/uptime-probe/internal/probe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatesMonitorWithName(t *testing.T) {
	monitor := NewMonitor("Some Monitor")
	assert.Equal(t, monitor.Name, "Some Monitor")
}

func TestReturnsThatMonitorIsUpWhenLatestResultIsSuccessful(t *testing.T) {
	monitor := Monitor{Name: "", historicalData: []probe.Result{
		{Status: probe.StatusSucceed},
	}}

	assert.Equal(t, Status(StatusUp), monitor.Status())
}

func TestReturnsDownWhenLastResultIsAFailure(t *testing.T) {
	monitor := Monitor{Name: "", historicalData: []probe.Result{
		{Status: probe.StatusFailed},
	}}

	assert.Equal(t, Status(StatusDown), monitor.Status())
}

func TestReturnsUnknownWhenNoProbeHasExecuted(t *testing.T) {
	monitor := NewMonitor("Some")

	assert.Equal(t, Status(StatusUnknown), monitor.Status())
}
