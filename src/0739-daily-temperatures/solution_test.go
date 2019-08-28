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
				[]int{},
			},
			[]interface{}{
				[]int{},
			},
		},
		{
			"Test",
			[]interface{}{
				[]int{1},
			},
			[]interface{}{
				[]int{0},
			},
		},
		{
			"Test",
			[]interface{}{
				[]int{73, 74, 75, 71, 69, 72, 76, 73},
			},
			[]interface{}{
				[]int{1, 1, 4, 2, 1, 1, 0, 0},
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := dailyTemperatures(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
