package solution

func letterCombinations(digits string) []string {
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	ret := []string{}
	if digits == "" {
		return ret
	}
	ret = m[string(digits[0])]
	if len(digits) == 1 {
		return ret
	}

	for i := 1; i < len(digits); i++ {
		newRet := []string{}
		for _, c1 := range ret {
			for _, c2 := range m[string(digits[i])] {
				newRet = append(newRet, c1+c2)
			}
		}
		ret = newRet
	}
	return ret
}
