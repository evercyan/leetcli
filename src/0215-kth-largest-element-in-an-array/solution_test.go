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
				[]int{3, 2, 1, 5, 6, 4},
				2,
			},
			[]interface{}{
				5,
			},
		},
		{
			[]interface{}{
				[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				4,
			},
			[]interface{}{
				4,
			},
		},
		{
			[]interface{}{
				[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				100,
			},
			[]interface{}{
				-1,
			},
		},
	}
	for _, c := range cases {
		t.Run("findKthLargest", func(t *testing.T) {
			ret := findKthLargest(c.inputs[0].([]int), c.inputs[1].(int))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
