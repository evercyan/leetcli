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
				[]int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			[]interface{}{
				4,
			},
		},
		{
			[]interface{}{
				[]int{0, 1, 0, 3, 2, 3},
			},
			[]interface{}{
				4,
			},
		},
		{
			[]interface{}{
				[]int{7, 7, 7, 7, 7, 7, 7},
			},
			[]interface{}{
				1,
			},
		},
	}
	for _, c := range cases {
		t.Run("lengthOfLIS", func(t *testing.T) {
			ret := lengthOfLIS(c.inputs[0].([]int))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
