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
				[]int{1, 2, 3},
			},
			[]interface{}{
				[][]int{
					{},
					{1},
					{2},
					{3},
					{1, 2},
					{1, 3},
					{2, 3},
					{1, 2, 3},
				},
			},
		},
		{
			[]interface{}{
				[]int{9, 0, 3, 5, 7},
			},
			[]interface{}{
				[][]int{
					{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run("subsets", func(t *testing.T) {
			ret := subsets(c.inputs[0].([]int))
			assert.ElementsMatch(t, c.expects[0].([][]int), ret)
		})
	}
}
