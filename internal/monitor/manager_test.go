package monitor

import (
	"context"
	"fmt"
	"github.com/RykoL/uptime-probe/config"
	"github.com/RykoL/uptime-probe/internal/monitor/probe"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

var httpProbe = probe.NewHttpProbe("http://localhost:8080")

func TestCreatesMonitorForEveryEntryInConfiguration(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080"},
	}}

	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)

	repo.EXPECT().SaveMonitor(gomock.Any(), gomock.Any()).Times(1)

	m := NewManager(logger, repo)
	m.applyConfig(&cfg)

	assert.Equal(t, 1, len(m.monitors))
}

func TestManager_Run_ReturnsErrorIfManagerIsNotInitialized(t *testing.T) {
	m := NewManager(nil, nil)

	err := m.Run(context.Background())

	assert.Error(t, err, fmt.Errorf("manager is not initialized yet. Call Initialize() before running"))
}

func TestManager_Init_SetsManagerToInitializedWhenSuccessful(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)

	repo.EXPECT().GetMonitors(gomock.Any()).Return(make([]*Monitor, 0), nil)

	m := NewManager(logger, repo)

	m.Initialize(context.Background(), &config.Config{})

	assert.True(t, m.initialized)
}

func TestManager_Init_DoesNotSetManagerToInitializedWhenFailing(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)

	repo.EXPECT().GetMonitors(gomock.Any()).Return(make([]*Monitor, 0), fmt.Errorf(""))

	m := NewManager(logger, repo)

	err := m.Initialize(context.Background(), &config.Config{})

	assert.Error(t, err)
	assert.False(t, m.initialized)
}

func TestManager_ApplyConfig_DoesNotAddMonitorFromConfigIfItAlreadyExistsInTheManager(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080", Interval: oneSecond},
	}}

	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)

	existingMonitor := []*Monitor{
		NewMonitor("TestMonitor", oneSecond, httpProbe),
	}

	repo.EXPECT().GetMonitors(gomock.Any()).Return(existingMonitor, nil)

	m := NewManager(logger, repo)

	m.Initialize(context.Background(), &cfg)

	assert.Len(t, m.monitors, 1)
}

func TestManager_ApplyConfig_DoesAddMonitorFromConfigIfMonitorDoesntExistYet(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080", Interval: oneSecond},
	}}

	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)

	existingMonitor := []*Monitor{
		NewMonitor("MyNewAndDifferentMonitor", oneSecond, httpProbe),
	}

	repo.EXPECT().GetMonitors(gomock.Any()).Return(existingMonitor, nil)
	repo.EXPECT().SaveMonitor(gomock.Any(), gomock.Any()).Times(1)

	m := NewManager(logger, repo)

	m.Initialize(context.Background(), &cfg)

	assert.Len(t, m.monitors, 2)
}

func TestManager_ApplyConfig_PersistsNewMonitors(t *testing.T) {
	cfg := config.Config{Monitors: []*config.MonitorConfig{
		{Name: "TestMonitor", Url: "http://localhost:8080", Interval: oneSecond},
	}}

	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)

	repo.EXPECT().SaveMonitor(gomock.Any(), gomock.Any()).Times(1)

	m := NewManager(logger, repo)

	m.applyConfig(&cfg)

}
