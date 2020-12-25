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
				"3[a]2[bc]",
			},
			[]interface{}{
				"aaabcbc",
			},
		},
		{
			[]interface{}{
				"3[a2[c]]",
			},
			[]interface{}{
				"accaccacc",
			},
		},
		{
			[]interface{}{
				"2[abc]3[cd]ef",
			},
			[]interface{}{
				"abcabccdcdcdef",
			},
		},
		{
			[]interface{}{
				"abc3[cd]xyz",
			},
			[]interface{}{
				"abccdcdcdxyz",
			},
		},
	}
	for _, c := range cases {
		t.Run("decodeString", func(t *testing.T) {
			ret := decodeString(c.inputs[0].(string))
			assert.Equal(t, c.expects[0].(string), ret)
		})
	}
}
