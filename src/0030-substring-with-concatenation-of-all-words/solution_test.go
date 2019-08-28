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
				"barfoothefoobarman",
				[]string{"foo", "bar"},
			},
			[]interface{}{
				[]int{0, 9},
			},
		},
		{
			"Test 2",
			[]interface{}{
				"wordgoodgoodgoodbestword",
				[]string{"word", "good", "best", "word"},
			},
			[]interface{}{
				[]int{},
			},
		},
		{
			"Test 3",
			[]interface{}{
				"barfoofoobarthefoobarman",
				[]string{"bar", "foo", "the"},
			},
			[]interface{}{
				[]int{6, 9, 12},
			},
		},
		{
			"Test 4",
			[]interface{}{
				"barfoofoobarthefoobarman",
				[]string{},
			},
			[]interface{}{
				[]int{},
			},
		},
		{
			"Test 5",
			[]interface{}{
				"wordgoodgoodgoodbestword",
				[]string{"word", "good", "best", "good"},
			},
			[]interface{}{
				[]int{8},
			},
		},
		{
			"Test 6",
			[]interface{}{
				"lingmindraboofooowingdingbarrwingmonkeypoundcake",
				[]string{"fooo", "barr", "wing", "ding", "wing"},
			},
			[]interface{}{
				[]int{13},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findSubstring(c.inputs[0].(string), c.inputs[1].([]string))
			if !reflect.DeepEqual(ret, c.expects[0].([]int)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
