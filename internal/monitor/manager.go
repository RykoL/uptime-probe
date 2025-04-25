package monitor

import (
	"context"
	"github.com/RykoL/uptime-probe/config"
	"github.com/RykoL/uptime-probe/internal/monitor/probe"
	"log/slog"
)

type Manager struct {
	monitors   []*Monitor
	log        *slog.Logger
	repository *Repository
}

func NewManager(logger *slog.Logger, repository *Repository) Manager {
	return Manager{log: logger, repository: repository}
}

func (m *Manager) ApplyConfig(cfg *config.Config) {
	for _, monitorConfig := range cfg.Monitors {
		target := monitorConfig.Url

		newProbe := probe.NewHttpProbe(target)
		m.monitors = append(m.monitors, NewMonitor(monitorConfig.Name, monitorConfig.Interval, newProbe))
	}
}

func (m *Manager) reconcile(ctx context.Context) {
	//existingMonitors, err := m.repository.GetMonitors(ctx)
}

func (m *Manager) Run() {
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
