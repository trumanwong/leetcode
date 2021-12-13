package countMatches

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountMatches(t *testing.T) {
	tests := []struct {
		items     [][]string
		ruleKey   string
		ruleValue string
		output    int
	}{
		{[][]string{{"phone","blue","pixel"},{"computer","silver","lenovo"},{"phone","gold","iphone"}}, "color", "silver", 1},{[][]string{{"phone","blue","pixel"},{"computer","silver","lenovo"},{"phone","gold","iphone"}}, "type", "phone", 2},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, countMatches(test.items, test.ruleKey, test.ruleValue))
	}
}
