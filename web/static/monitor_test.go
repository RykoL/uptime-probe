package static

import (
	"github.com/RykoL/uptime-probe/internal/testutils"
	"github.com/RykoL/uptime-probe/web/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRendersMonitorName(t *testing.T) {
	m := model.Monitor{Name: "SomeMonitor"}
	comp := testutils.RenderComponent(Monitor(m))

	assert.Equal(t, "SomeMonitor", comp.Find(`h2`).Text())
}

type statusCase struct {
	Text string
}

func TestRendersStatus(t *testing.T) {
	cases := []statusCase{
		{"Down"},
		{"Up"},
		{"Unknown"},
		{"Pending"},
	}

	for _, testCase := range cases {
		comp := testutils.RenderComponent(MonitorStatus(testCase.Text))

		assert.Equal(t, testCase.Text, comp.Find(`span`).Text())

	}

}
