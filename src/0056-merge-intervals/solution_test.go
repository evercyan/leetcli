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
				[][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			[]interface{}{
				[][]int{
					{1, 6},
					{8, 10},
					{15, 18},
				},
			},
		},
		{
			[]interface{}{
				[][]int{
					{1, 4},
					{4, 5},
				},
			},
			[]interface{}{
				[][]int{
					{1, 5},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run("merge", func(t *testing.T) {
			ret := merge(c.inputs[0].([][]int))
			assert.Equal(t, c.expects[0].([][]int), ret)
		})
	}
}
