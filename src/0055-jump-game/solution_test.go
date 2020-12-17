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
				[]int{2, 3, 1, 1, 4},
			},
			[]interface{}{
				true,
			},
		},
		{
			[]interface{}{
				[]int{3, 2, 1, 0, 4},
			},
			[]interface{}{
				false,
			},
		},
		{
			[]interface{}{
				[]int{0},
			},
			[]interface{}{
				true,
			},
		},
		{
			[]interface{}{
				[]int{2, 0, 0},
			},
			[]interface{}{
				true,
			},
		},
	}
	for _, c := range cases {
		t.Run("canJump", func(t *testing.T) {
			ret := canJump(c.inputs[0].([]int))
			assert.Equal(t, c.expects[0].(bool), ret)
		})
	}
}
