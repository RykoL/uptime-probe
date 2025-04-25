package monitor

import (
	"fmt"
	"github.com/RykoL/uptime-probe/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatesMonitorForEveryEntryInConfiguration(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080"},
	}}

	m := NewManager(nil, nil)
	m.ApplyConfig(&cfg)

	assert.Equal(t, 1, len(m.monitors))
}

func TestManager_Run_ReturnsErrorIfManagerIsNotInitialized(t *testing.T) {
	m := NewManager(nil, nil)

	err := m.Run()

	assert.Error(t, err, fmt.Errorf("manager is not initialized yet. Call Init() before running"))
}
