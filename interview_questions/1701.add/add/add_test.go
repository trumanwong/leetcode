package add

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		output int
	}{
		{1, 2, 3},
		{5, 6, 11},
	}

	for _, test := range tests {
		ret := add(test.a, test.b)
		if ret != test.output {
			t.Errorf("Got %d for input (%d+%d); expected %d", ret, test.a, test.b, test.output)
		}
	}
}
