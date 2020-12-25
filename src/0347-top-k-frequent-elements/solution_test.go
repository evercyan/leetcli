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
				[]int{1, 1, 1, 2, 2, 3},
				2,
			},
			[]interface{}{
				[]int{1, 2},
			},
		},
		{
			[]interface{}{
				[]int{1},
				1,
			},
			[]interface{}{
				[]int{1},
			},
		},
	}
	for _, c := range cases {
		t.Run("topKFrequent", func(t *testing.T) {
			ret := topKFrequent(c.inputs[0].([]int), c.inputs[1].(int))
			assert.Equal(t, c.expects[0].([]int), ret)
		})
	}
}
