package decode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		encoded  []int
		first    int
		excepted []int
	}{
		{[]int{1, 2, 3}, 1, []int{1, 0, 2, 1}},
		{[]int{6, 2, 7, 3}, 4, []int{4, 2, 0, 7, 4}},
	}

	for _, test := range tests {
		assert.Equal(t, test.excepted, decode(test.encoded, test.first))
	}
}
