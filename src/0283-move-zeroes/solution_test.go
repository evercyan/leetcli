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
				[]int{0, 1, 0, 3, 12},
			},
			[]interface{}{
				[]int{1, 3, 12, 0, 0},
			},
		},
	}
	for _, c := range cases {
		t.Run("moveZeroes", func(t *testing.T) {
			nums := c.inputs[0].([]int)
			moveZeroes(nums)
			assert.Equal(t, c.expects[0].([]int), nums)
		})
	}
}
