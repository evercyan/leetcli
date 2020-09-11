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
				[]int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			[]interface{}{
				[]int{5, 6},
			},
		},
		{
			[]interface{}{
				[]int{},
			},
			[]interface{}{
				[]int{},
			},
		},
	}
	for _, c := range cases {
		t.Run("findDisappearedNumbers", func(t *testing.T) {
			ret := findDisappearedNumbers(c.inputs[0].([]int))
			assert.ElementsMatch(t, c.expects[0].([]int), ret)
		})
	}
}
