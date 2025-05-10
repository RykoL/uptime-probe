package monitor

import (
	"context"
	"fmt"
	"github.com/RykoL/uptime-probe/config"
	"github.com/RykoL/uptime-probe/internal/monitor/probe"
	"log/slog"
	"sync"
)

type Manager struct {
	monitors    []*Monitor
	log         *slog.Logger
	repository  Repository
	initialized bool
	wg          *sync.WaitGroup
}

func NewManager(logger *slog.Logger, repository Repository) Manager {
	return Manager{log: logger, repository: repository}
}

func (m *Manager) Initialize(ctx context.Context, cfg *config.Config) error {
	monitors, err := m.repository.GetMonitors(ctx)

	if err != nil {
		return err
	}

	m.monitors = monitors
	m.applyConfig(cfg)
	m.log.Info("Monitors loaded", slog.Int("count", len(m.monitors)))
	m.initialized = true
	return nil
}

func (m *Manager) applyConfig(cfg *config.Config) {
	for _, monitorConfig := range cfg.Monitors {
		newMonitor := NewMonitor(monitorConfig.Name, monitorConfig.Interval, probe.NewHttpProbe(monitorConfig.Url))
		if !m.monitorExists(newMonitor) {
			m.log.Info("Found new monitor from config", "name", newMonitor.Name)
			id, err := m.repository.SaveMonitor(context.Background(), newMonitor)
			if err != nil {
				m.log.Error("Failed to persist monitor", "name", newMonitor.Name, "error", err)
			}
			newMonitor.Id = id
			m.monitors = append(m.monitors, newMonitor)
		}
	}

}

func (m *Manager) monitorExists(monitor *Monitor) bool {
	for _, existingMonitor := range m.monitors {
		if existingMonitor.Equals(monitor) {
			return true
		}
	}

	return false
}

func (m *Manager) Run(ctx context.Context) error {

	if !m.initialized {
		return fmt.Errorf("manager is not initialized yet. Call Initialize() before running")
	}

	var wg sync.WaitGroup

	for _, monitor := range m.monitors {

		wg.Add(1)

		go func() {
			defer wg.Done()
			monitor.Start(ctx, m.repository)
		}()
	}

	return nil
}

func (m *Manager) Stop() {
	m.wg.Wait()
}
