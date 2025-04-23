package config_test

import (
	"github.com/RykoL/uptime-probe/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadsSimpleConfiguration(t *testing.T) {
	cfg, err := config.LoadFromFile("testdata/example.yaml")

	assert.NoError(t, err)
	assert.Equal(t, 2, len(cfg.Monitors))
}
