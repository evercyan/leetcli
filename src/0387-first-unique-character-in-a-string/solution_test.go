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
				"leetcode",
			},
			[]interface{}{
				0,
			},
		},
		{
			[]interface{}{
				"loveleetcode",
			},
			[]interface{}{
				2,
			},
		},
		{
			[]interface{}{
				"lll",
			},
			[]interface{}{
				-1,
			},
		},
	}
	for _, c := range cases {
		t.Run("firstUniqChar", func(t *testing.T) {
			ret := firstUniqChar(c.inputs[0].(string))
			assert.Equal(t, c.expects[0].(int), ret)
		})
	}
}
