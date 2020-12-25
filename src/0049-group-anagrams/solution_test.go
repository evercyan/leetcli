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
				[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			[]interface{}{
				[][]string{
					{"eat", "tea", "ate"},
					{"tan", "nat"},
					{"bat"},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run("groupAnagrams", func(t *testing.T) {
			ret := groupAnagrams(c.inputs[0].([]string))
			assert.ElementsMatch(t, c.expects[0].([][]string), ret)
		})
	}
}
