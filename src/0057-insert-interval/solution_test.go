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
				[]Interval{
					Interval{
						Start: 1,
						End:   3,
					},
					Interval{
						Start: 6,
						End:   9,
					},
				},
				Interval{
					Start: 2,
					End:   5,
				},
			},
			[]interface{}{
				[]Interval{
					Interval{
						Start: 1,
						End:   5,
					},
					Interval{
						Start: 6,
						End:   9,
					},
				},
			},
		},
		{
			"Test 2",
			[]interface{}{
				[]Interval{
					Interval{
						Start: 1,
						End:   2,
					},
					Interval{
						Start: 3,
						End:   5,
					},
					Interval{
						Start: 6,
						End:   7,
					},
					Interval{
						Start: 8,
						End:   10,
					},
					Interval{
						Start: 12,
						End:   16,
					},
				},
				Interval{
					Start: 4,
					End:   8,
				},
			},
			[]interface{}{
				[]Interval{
					Interval{
						Start: 1,
						End:   2,
					},
					Interval{
						Start: 3,
						End:   10,
					},
					Interval{
						Start: 12,
						End:   16,
					},
				},
			},
		},
		{
			"Test 3",
			[]interface{}{
				[]Interval{},
				Interval{
					Start: 4,
					End:   8,
				},
			},
			[]interface{}{
				[]Interval{
					Interval{
						Start: 4,
						End:   8,
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := insert(c.inputs[0].([]Interval), c.inputs[1].(Interval))
			if !reflect.DeepEqual(ret, c.expects[0].([]Interval)) {
				t.Fatalf("FAIL ----> name: %v, inputs: %v, expects: %v, ret: %v", c.name, c.inputs, c.expects[0], ret)
			}
		})
	}
}
