package decode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		encoded  []int
		excepted []int
	}{
		{[]int{3, 1}, []int{1, 2, 3}},
		{[]int{6, 5, 4, 6}, []int{2, 4, 1, 5, 3}},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, decode(test.encoded))
	}
}
