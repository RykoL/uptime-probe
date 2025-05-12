package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMonitor_OldestProbeResult(t *testing.T) {
	oneMinute, _ := time.ParseDuration("1m")
	m := Monitor{Results: []ProbeResult{
		{
			Timestamp: time.Now().Add(-2 * oneMinute),
			Success:   false,
		},
		{
			Timestamp: time.Now().Add(-oneMinute),
			Success:   false,
		},
		{
			Timestamp: time.Now(),
			Success:   false,
		},
	}}

	expected := m.Results[0]
	assert.Equal(t, expected, m.OldestProbeResult())
}

func TestMonitor_LatestProbeResult(t *testing.T) {
	oneMinute, _ := time.ParseDuration("1m")
	m := Monitor{Results: []ProbeResult{
		{
			Timestamp: time.Now().Add(-2 * oneMinute),
			Success:   false,
		},
		{
			Timestamp: time.Now().Add(-oneMinute),
			Success:   false,
		},
		{
			Timestamp: time.Now(),
			Success:   false,
		},
	}}

	expected := m.Results[2]
	assert.Equal(t, expected, m.LatestProbeResult())
}
