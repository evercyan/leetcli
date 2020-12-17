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
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			[]interface{}{
				7,
			},
		},
	}

	for _, c := range cases {
		t.Run("minPathSum", func(t *testing.T) {
			ret := minPathSum(c.inputs[0].([][]int))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
