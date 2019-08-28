package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []interface{}
		expects []interface{}
	}{
		{
			"Test 1",
			[]interface{}{
				"23",
			},
			[]interface{}{
				[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
			},
		},
		{
			"Test 2",
			[]interface{}{
				"2",
			},
			[]interface{}{
				[]string{"a", "b", "c"},
			},
		},
		{
			"Test 3",
			[]interface{}{
				"",
			},
			[]interface{}{
				[]string{},
			},
		},
		{
			"Test 4",
			[]interface{}{
				"234",
			},
			[]interface{}{
				[]string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := letterCombinations(c.inputs[0].(string))
			if len(ret) != len(c.expects[0].([]string)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
