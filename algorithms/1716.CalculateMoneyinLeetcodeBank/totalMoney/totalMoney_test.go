package totalMoney

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalMoney(t *testing.T) {
	tests := []struct {
		n        int
		excepted int
	}{
		{4, 10},
		{10, 37},
		{20, 96},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, totalMoney(test.n))
	}
}
