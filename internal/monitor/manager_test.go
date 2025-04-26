package monitor

import (
	"context"
	"fmt"
	"github.com/RykoL/uptime-probe/config"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

type RepositorySpy struct {
	shouldReturnError bool
}

func (r *RepositorySpy) GetMonitors(ctx context.Context) ([]*Monitor, error) {
	if r.shouldReturnError {
		return nil, fmt.Errorf("")
	}
	return make([]*Monitor, 0), nil
}

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

	assert.Error(t, err, fmt.Errorf("manager is not initialized yet. Call Initialize() before running"))
}

func TestManager_Init_SetsManagerToInitializedWhenSuccessful(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	m := NewManager(logger, &RepositorySpy{})

	m.Initialize(context.Background())

	assert.True(t, m.initialized)
}

func TestManager_Init_DoesNotSetManagerToInitializedWhenFailing(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	m := NewManager(logger, &RepositorySpy{shouldReturnError: true})

	err := m.Initialize(context.Background())

	assert.Error(t, err)
	assert.False(t, m.initialized)
}
