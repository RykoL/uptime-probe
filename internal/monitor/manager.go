package monitor

import (
	"context"
	"fmt"
	"github.com/RykoL/uptime-probe/config"
	"github.com/RykoL/uptime-probe/internal/monitor/probe"
	"log/slog"
)

type Manager struct {
	monitors    []*Monitor
	log         *slog.Logger
	repository  Repository
	initialized bool
}

func NewManager(logger *slog.Logger, repository Repository) Manager {
	return Manager{log: logger, repository: repository}
}

func (m *Manager) Initialize(ctx context.Context) error {
	monitors, err := m.repository.GetMonitors(ctx)

	if err != nil {
		return err
	}

	m.monitors = monitors
	m.log.Info("Monitors loaded", slog.Int("count", len(monitors)))
	m.initialized = true
	return nil
}

func (m *Manager) ApplyConfig(cfg *config.Config) {
	for _, monitorConfig := range cfg.Monitors {
		target := monitorConfig.Url

		newProbe := probe.NewHttpProbe(target)
		m.monitors = append(m.monitors, NewMonitor(monitorConfig.Name, monitorConfig.Interval, newProbe))
	}
}

/*func (m *Manager) reconcile(ctx context.Context) {
	existingMonitors, err := m.repository.GetMonitors(ctx)
}*/

func (m *Manager) Run() error {

	if !m.initialized {
		return fmt.Errorf("manager is not initialized yet. Call Initialize() before running")
	}

	for {
		for _, monitor := range m.monitors {

			if !monitor.ShouldExecuteProbe() {
				continue
			}

			m.log.Info("Executing probe", "monitor_name", monitor.Name)
			if err := monitor.Probe(); err != nil {
				m.log.Warn(err.Error(), "monitor_name", monitor.Name)
			}
		}
	}
}
