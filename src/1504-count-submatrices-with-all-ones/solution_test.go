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
				[][]int{
					{1, 0, 1},
					{1, 1, 0},
					{1, 1, 0},
				},
			},
			[]interface{}{
				13,
			},
		},
		{
			"Test",
			[]interface{}{
				[][]int{
					{0, 1, 1, 0},
					{0, 1, 1, 1},
					{1, 1, 1, 0},
				},
			},
			[]interface{}{
				24,
			},
		},
		{
			"Test",
			[]interface{}{
				[][]int{
					{1, 1, 1, 1, 1, 1},
				},
			},
			[]interface{}{
				21,
			},
		},
		{
			"Test",
			[]interface{}{
				[][]int{
					{1, 0, 1},
					{0, 1, 0},
					{1, 0, 1},
				},
			},
			[]interface{}{
				5,
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := numSubmat(c.inputs[0].([][]int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}
