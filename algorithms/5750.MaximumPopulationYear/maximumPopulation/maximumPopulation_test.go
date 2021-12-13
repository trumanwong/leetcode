package maximumPopulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumPopulation(t *testing.T) {
	tests := []struct {
		logs     [][]int
		excepted int
	}{
		{[][]int{{1993, 1999}, {2000, 2010}}, 1993},
		{[][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}, 1960},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, maximumPopulation(test.logs))
	}
}
