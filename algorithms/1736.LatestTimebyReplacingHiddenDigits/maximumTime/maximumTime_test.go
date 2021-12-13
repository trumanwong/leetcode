package maximumTime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumTime(t *testing.T) {
	tests := []struct {
		time     string
		excepted string
	}{
		{"2?:?0", "23:50"},
		{"0?:3?", "09:39"},
		{"1?:22", "19:22"},
		{"??:3?", "23:39"},
		{"??:??", "23:59"},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, maximumTime(test.time))
	}
}
