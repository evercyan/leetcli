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
				[]int{3, 2, 3},
			},
			[]interface{}{
				3,
			},
		},
		{
			[]interface{}{
				[]int{2, 2, 1, 1, 1, 2, 2},
			},
			[]interface{}{
				2,
			},
		},
		{
			[]interface{}{
				[]int{6, 5, 5},
			},
			[]interface{}{
				5,
			},
		},
	}
	for _, c := range cases {
		t.Run("majorityElement", func(t *testing.T) {
			ret := majorityElement(c.inputs[0].([]int))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
