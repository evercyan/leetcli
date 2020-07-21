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
				1, 2, 3,
			},
			[]interface{}{
				[]int{3, 4, 5, 6},
			},
		},
		{
			"Test",
			[]interface{}{
				1, 1, 1,
			},
			[]interface{}{
				[]int{1},
			},
		},
		{
			"Test",
			[]interface{}{
				1, 2, 0,
			},
			[]interface{}{
				[]int{},
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := divingBoard(c.inputs[0].(int), c.inputs[1].(int), c.inputs[2].(int))
			if !reflect.DeepEqual(ret, c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
