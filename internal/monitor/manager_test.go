package monitor

import (
	"context"
	"fmt"
	"github.com/RykoL/uptime-probe/config"
	"github.com/RykoL/uptime-probe/internal/monitor/probe"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
)

type RepositorySpy struct {
	monitors          []*Monitor
	shouldReturnError bool
	CalledSave        bool
}

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func (r *RepositorySpy) GetMonitors(ctx context.Context) ([]*Monitor, error) {
	if r.shouldReturnError {
		return nil, fmt.Errorf("")
	}
	return r.monitors, nil
}

func (r *RepositorySpy) SaveMonitor(ctx context.Context, monitor *Monitor) error {
	r.CalledSave = true
	return nil
}

func TestCreatesMonitorForEveryEntryInConfiguration(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080"},
	}}

	m := NewManager(logger, &RepositorySpy{})
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
	m := NewManager(logger, &RepositorySpy{shouldReturnError: true})

	err := m.Initialize(context.Background())

	assert.Error(t, err)
	assert.False(t, m.initialized)
}

func TestManager_ApplyConfig_DoesNotAddMonitorFromConfigIfItAlreadyExistsInTheManager(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080", Interval: oneSecond},
	}}

	httpProbe := probe.NewHttpProbe("http://localhost:8080")

	m := NewManager(logger, &RepositorySpy{
		shouldReturnError: false,
		monitors: []*Monitor{
			NewMonitor("TestMonitor", oneSecond, httpProbe),
		},
	})

	m.Initialize(context.Background())

	assert.Len(t, m.monitors, 1)

	m.ApplyConfig(&cfg)

	assert.Len(t, m.monitors, 1)
}

func TestManager_ApplyConfig_DoesAddMonitorFromConfigIfMonitorDoesntExistYet(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080", Interval: oneSecond},
	}}

	httpProbe := probe.NewHttpProbe("http://localhost:3000")

	m := NewManager(logger, &RepositorySpy{
		shouldReturnError: false,
		monitors: []*Monitor{
			NewMonitor("MyNewAndDifferentMonitor", oneSecond, httpProbe),
		},
	})

	m.Initialize(context.Background())

	assert.Len(t, m.monitors, 1)

	m.ApplyConfig(&cfg)

	assert.Len(t, m.monitors, 2)
}

func TestManager_ApplyConfig_PersistsNewMonitors(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080", Interval: oneSecond},
	}}
	repo := &RepositorySpy{
		shouldReturnError: false,
		monitors:          make([]*Monitor, 0),
	}
	m := NewManager(logger, repo)

	m.ApplyConfig(&cfg)

	assert.True(t, repo.CalledSave)
}
