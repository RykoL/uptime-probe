package static

import (
	"github.com/RykoL/uptime-probe/internal/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeader(t *testing.T) {
	// Pipe the rendered template into goquery.
	doc := testutils.RenderComponent(Index())
	assert.NotNil(t, doc.Find(`h1[aria-label="Monitors"]`))
}

func TestShowDefaultMessageIfNoMonitorsExist(t *testing.T) {
	doc := testutils.RenderComponent(Index())
	assert.Equal(t, "No monitors created yet.", doc.Find(`p`).Text())
}
