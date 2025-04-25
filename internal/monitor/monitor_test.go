package monitor

import (
	"github.com/RykoL/uptime-probe/internal/monitor/probe"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type NoOpProbe struct {
}

var oneSecond, _ = time.ParseDuration("1s")
var oneMinute, _ = time.ParseDuration("1m")

func (p *NoOpProbe) Target() string {
	return ""
}

func (p *NoOpProbe) Execute() (*probe.ProbeResult, error) {
	return &probe.ProbeResult{Succeeded: true}, nil
}

func TestCreatesMonitorWithName(t *testing.T) {
	monitor := NewMonitor("Some Monitor", oneMinute, &NoOpProbe{})
	assert.Equal(t, monitor.Name, "Some Monitor")
}

func TestReturnsThatMonitorIsUpWhenLatestResultIsSuccessful(t *testing.T) {
	monitor := Monitor{Name: "", previousProbes: []*probe.ProbeResult{
		{Succeeded: probe.ExecutionSucceeded},
	}}

	assert.Equal(t, Status(StatusUp), monitor.Status())
}

func TestReturnsDownWhenLastResultIsAFailure(t *testing.T) {
	monitor := Monitor{Name: "", previousProbes: []*probe.ProbeResult{
		{Succeeded: probe.ExecutionFailed},
	}}

	assert.Equal(t, Status(StatusDown), monitor.Status())
}

func TestReturnsUnknownWhenNoProbeHasExecuted(t *testing.T) {
	monitor := NewMonitor("Some", oneMinute, &NoOpProbe{})

	assert.Equal(t, Status(StatusUnknown), monitor.Status())
}

func TestReturnsTrueIfMonitorHasNotExecutedAProbeYet(t *testing.T) {
	m := NewMonitor("", oneMinute, &NoOpProbe{})

	assert.Equal(t, true, m.ShouldExecuteProbe())
}

func TestReturnsTrueIfLastExecutionLiesBehindInterval(t *testing.T) {
	lastExecution := time.Now().Add(-oneMinute)
	m := Monitor{
		Name:     "asdasd",
		Interval: oneSecond,
		previousProbes: []*probe.ProbeResult{
			{Succeeded: true, TimeStamp: lastExecution},
		},
	}

	assert.Equal(t, true, m.ShouldExecuteProbe())
}

func TestReturnsFalseIfLastExecutionLiesAfterInterval(t *testing.T) {
	lastExecution := time.Now().Add(-oneSecond)
	m := Monitor{
		Name:     "asdasd",
		Interval: oneMinute,
		previousProbes: []*probe.ProbeResult{
			{Succeeded: true, TimeStamp: lastExecution},
		},
	}

	assert.Equal(t, false, m.ShouldExecuteProbe())
}

func TestMonitor_Probe_StoresProbeResult(t *testing.T) {
	m := NewMonitor("", oneMinute, &NoOpProbe{})

	assert.Empty(t, m.GetPreviousProbes())

	m.Probe()

	assert.NotEmpty(t, m.GetPreviousProbes())
}

func TestMonitor_ShouldExecuteProbe_ReturnsFalseDirectlyAfterExecuting(t *testing.T) {
	m := NewMonitor("", oneMinute, &NoOpProbe{})

	assert.True(t, m.ShouldExecuteProbe())

	m.Probe()

	assert.False(t, m.ShouldExecuteProbe())
}

func TestNewMonitorFromRecord_CorrectlyMapsFields(t *testing.T) {
	record := monitorRecord{
		Id:         0,
		Name:       "MyMonitor",
		Interval:   oneMinute,
		Definition: `{"url": "https: //google.com"}`,
	}

	m, err := NewMonitorFromRecord(record)

	assert.NoError(t, err)
	assert.Equal(t, "MyMonitor", m.Name)
	assert.Equal(t, oneMinute, m.Interval)
	assert.NotNil(t, m.probe)
}

func TestIsSameAs_ReturnsTrueWhenBothMonitorsShareSameAttributes(t *testing.T) {
	m1 := NewMonitor("FirstMonitor", oneMinute, probe.NewHttpProbe("https://api.mytest.foo"))
	m2 := NewMonitor("FirstMonitor", oneMinute, probe.NewHttpProbe("https://api.mytest.foo"))

	assert.True(t, m1.IsSameAs(m2))
}

func TestIsSameAs_ReturnsFalseWhenBothMonitorsHaveDifferentAttributes(t *testing.T) {
	m1 := NewMonitor("FirstMonitor", oneMinute, probe.NewHttpProbe("https://api.mytest.foo"))
	m2 := NewMonitor("SecondMonitor", oneMinute, probe.NewHttpProbe("https://api.mytest.foo"))

	assert.False(t, m1.IsSameAs(m2))
}
