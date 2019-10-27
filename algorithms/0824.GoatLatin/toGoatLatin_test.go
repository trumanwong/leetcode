package main

import (
	"leetcode/algorithms/0824.GoatLatin/toGoatLatin"
	"testing"
)

func TestToGoatLatin(t *testing.T)  {
	tests := []struct{
		S string
		output string
	}{
		{"I speak Goat Latin","Imaa peaksmaaa oatGmaaaa atinLmaaaaa"},
		{"The quick brown fox jumped over the lazy dog",
			"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"},
	}

	for _, test := range tests {
		ret := toGoatLatin.ToGoatLatin(test.S)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.S, test.output)
		}
	}

}