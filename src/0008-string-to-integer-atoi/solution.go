package solution

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func stringToInt(v string) int {
	r, _ := strconv.Atoi(v)
	return r
}

func checkNum(sign int, result int) (int, bool) {
	result *= sign
	if result > math.MaxInt32 {
		return math.MaxInt32, false
	} else if result < math.MinInt32 {
		return math.MinInt32, false
	}
	return result, true
}

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	// 符号
	sign := 1
	if string(str[0]) == "-" {
		sign = -1
		str = str[1:]
	} else if string(str[0]) == "+" {
		str = str[1:]
	}
	// 结果
	result := 0
	for i := 0; i < len(str); i++ {
		matchRet, _ := regexp.Match(`^\d$`, []byte(string(str[i])))
		if matchRet {
			result = result*10 + stringToInt(string(str[i]))
			result, ret := checkNum(sign, result)
			if !ret {
				return result
			}
		} else {
			break
		}
	}
	result *= sign
	return result
}
