package config_test

import (
	"github.com/RykoL/uptime-probe/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadsSimpleConfiguration(t *testing.T) {
	cfg, err := config.LoadFromFile("testdata/example.yaml")

	assert.NoError(t, err)
	assert.Equal(t, 2, len(cfg.Monitors))
}

func TestLoadConfig_Unmarshalls_Monitor(t *testing.T) {
	cfg, err := config.LoadFromFile("testdata/singleMonitor.yaml")

	thirtySeconds, _ := time.ParseDuration("30s")

	assert.NoError(t, err)
	assert.Equal(t, "MyMonitor", cfg.Monitors[0].Name)
	assert.Equal(t, thirtySeconds, cfg.Monitors[0].Interval)
	assert.Equal(t, "https://google.com", cfg.Monitors[0].Url)
}
