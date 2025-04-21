package static

import (
	"github.com/RykoL/uptime-probe/internal/monitor"
	"github.com/RykoL/uptime-probe/internal/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRendersMonitorName(t *testing.T) {
	m := monitor.NewMonitor("SomeMonitor")
	comp := testutils.RenderComponent(Monitor(*m))

	assert.Equal(t, "SomeMonitor", comp.Find(`h2`).Text())
}

type statusCase struct {
	Status monitor.Status
	Text   string
}

func TestRendersStatus(t *testing.T) {
	cases := []statusCase{
		{monitor.StatusDown, "Down"},
		{monitor.StatusUp, "Up"},
		{monitor.StatusUnknown, "Unknown"},
		{monitor.StatusPending, "Pending"},
	}

	for _, testCase := range cases {
		comp := testutils.RenderComponent(MonitorStatus(testCase.Status))

		assert.Equal(t, testCase.Text, comp.Find(`span`).Text())

	}

}
