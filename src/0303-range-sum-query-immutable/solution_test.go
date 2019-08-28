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
			"Test 1",
			[]interface{}{
				0,
				2,
			},
			[]interface{}{
				1,
			},
		},
		{
			"Test 2",
			[]interface{}{
				2,
				5,
			},
			[]interface{}{
				-1,
			},
		},
		{
			"Test 3",
			[]interface{}{
				0,
				5,
			},
			[]interface{}{
				-3,
			},
		},
		{
			"Test 4",
			[]interface{}{
				0,
				0,
			},
			[]interface{}{
				-2,
			},
		},
		{
			"Test 5",
			[]interface{}{
				-1,
				0,
			},
			[]interface{}{
				0,
			},
		},
		{
			"Test 6",
			[]interface{}{
				0,
				100,
			},
			[]interface{}{
				0,
			},
		},
	}
	nums := []int{-2, 0, 3, -5, 2, -1}
	na := Constructor(nums)
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := na.SumRange(c.inputs[0].(int), c.inputs[1].(int))
			if !reflect.DeepEqual(ret, c.expects[0].(int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}

	// 单纯为过测试用例
	Constructor([]int{})
}
