package static

import (
	"github.com/RykoL/uptime-probe/internal/monitor"
	"github.com/RykoL/uptime-probe/internal/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var m monitor.Monitor = monitor.NewMonitor("My Monitor")

func TestHeader(t *testing.T) {
	// Pipe the rendered template into goquery.
	doc := testutils.RenderComponent(Index([]monitor.Monitor{m}))
	assert.NotNil(t, doc.Find(`h1[aria-label="Monitors"]`))
}

func TestShowDefaultMessageIfNoMonitorsExist(t *testing.T) {
	doc := testutils.RenderComponent(Index([]monitor.Monitor{}))
	assert.Equal(t, "No monitors created yet.", doc.Find(`p`).Text())
}

func TestShouldNotShowDefaultMessageIfMonitorsExists(t *testing.T) {
	doc := testutils.RenderComponent(Index([]monitor.Monitor{m}))

	assert.NotEqual(t, "No monitors created yet.", doc.Find(`p`).Text())
	assert.Equal(t, "My Monitor", doc.Find("h2").Text())
}
