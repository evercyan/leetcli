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
				[]int{1, 2, 3, 1},
			},
			[]interface{}{
				4,
			},
		},
		{
			"Test",
			[]interface{}{
				[]int{2, 7, 9, 3, 1},
			},
			[]interface{}{
				12,
			},
		},
		{
			"Test",
			[]interface{}{
				[]int{2, 1, 4, 5, 3, 1, 1, 3},
			},
			[]interface{}{
				12,
			},
		},
		{
			"Test",
			[]interface{}{
				[]int{},
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test",
			[]interface{}{
				[]int{1},
			},
			[]interface{}{
				1,
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := massage(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}

			ret2 := massage2(c.inputs[0].([]int))
			if !reflect.DeepEqual(ret2, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret2: %v", c.name, index, c.inputs, c.expects[0], ret2)
			}
		})
		index++
	}
}
