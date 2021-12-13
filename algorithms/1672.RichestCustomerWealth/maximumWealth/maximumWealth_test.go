package maximumWealth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumWealth(t *testing.T) {
	tests := []struct {
		accounts [][]int
		excepted int
	}{
		{[][]int{{1, 2, 3}, {3, 2, 1}}, 6},
		{[][]int{{1, 5}, {7, 3}, {3, 5}}, 10},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, maximumWealth(test.accounts))
	}
}
