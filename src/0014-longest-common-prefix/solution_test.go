package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []interface{}
		expects []interface{}
	}{
		{
			"Test",
			[]interface{}{
				[]string{"flower", "flow", "flight"},
			},
			[]interface{}{
				"fl",
			},
		},
		{
			"Test",
			[]interface{}{
				[]string{"dog", "racecar", "car"},
			},
			[]interface{}{
				"",
			},
		},
		{
			"Test",
			[]interface{}{
				[]string{},
			},
			[]interface{}{
				"",
			},
		},
		{
			"Test",
			[]interface{}{
				[]string{"hello"},
			},
			[]interface{}{
				"hello",
			},
		},
		{
			"Test",
			[]interface{}{
				[]string{"abca", "aba", "aaab"},
			},
			[]interface{}{
				"a",
			},
		},
		{
			"Test",
			[]interface{}{
				[]string{"c", "c"},
			},
			[]interface{}{
				"c",
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestCommonPrefix(c.inputs[0].([]string))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
