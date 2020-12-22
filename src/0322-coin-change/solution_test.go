package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		inputs  []interface{}
		expects []interface{}
	}{
		{
			[]interface{}{
				[]int{1, 2, 5},
				11,
			},
			[]interface{}{
				3,
			},
		},
		{
			[]interface{}{
				[]int{2},
				3,
			},
			[]interface{}{
				-1,
			},
		},
		{
			[]interface{}{
				[]int{1},
				0,
			},
			[]interface{}{
				0,
			},
		},
		{
			[]interface{}{
				[]int{1},
				2,
			},
			[]interface{}{
				2,
			},
		},
	}
	for _, c := range cases {
		t.Run("coinChange", func(t *testing.T) {
			ret := coinChange(c.inputs[0].([]int), c.inputs[1].(int))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
