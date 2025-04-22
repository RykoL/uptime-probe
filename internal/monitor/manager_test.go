package monitor

import (
	"github.com/RykoL/uptime-probe/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatesMonitorForEveryEntryInConfiguration(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080"},
	}}

	m := NewManager(nil)
	m.ApplyConfig(&cfg)

	assert.Equal(t, 1, len(m.monitors))
}
