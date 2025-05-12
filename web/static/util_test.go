package static

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatRelativeTime_FormatsMinuteCorrectly(t *testing.T) {

	now := time.Now()
	before := time.Now().Add(-time.Minute + time.Second)

	assert.Equal(t, "1m ago", FormatRelativeTime(now, before))
}
