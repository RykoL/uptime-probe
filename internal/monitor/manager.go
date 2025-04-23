package monitor

import (
	"github.com/RykoL/uptime-probe/config"
	"log/slog"
	"net/url"
	"time"
)

type Manager struct {
	monitors []*Monitor
	log      *slog.Logger
}

func NewManager(logger *slog.Logger) Manager {
	return Manager{log: logger}
}

func (m *Manager) ApplyConfig(cfg *config.Config) {
	for _, monitorConfig := range cfg.Monitors {
		target, err := url.Parse(monitorConfig.Url)

		if err != nil {
			m.log.Warn("Failed to create monitor %s: %v", monitorConfig.Name, err)
		}

		probe := NewHttpProbe(target)
		// TODO: Load interval from configuration
		interval, _ := time.ParseDuration("1m")
		m.monitors = append(m.monitors, NewMonitor(monitorConfig.Name, interval, &probe))
	}
}

func (m *Manager) Run() {
	for {
		for _, monitor := range m.monitors {
			m.log.Info("Executing probe", "monitor_name", monitor.Name)

			if err := monitor.Probe(); err != nil {
				m.log.Info("Failed executing probe", "monitor_name", monitor.Name)
			}
		}
	}
}
