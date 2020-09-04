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
				"abc",
				[]int{},
			},
			[]interface{}{
				"abc",
			},
		},
		{
			"Test",
			[]interface{}{
				"abc",
				[]int{3, 5, 9},
			},
			[]interface{}{
				"rpl",
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := shiftingLetters(c.inputs[0].(string), c.inputs[1].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
