package solution

import (
	"strconv"
)

/**
 * 题目解读
 * 数列第 1 项默认为 1, 后项以前项数进行报项
 * 例:
 * 第 2 项, 针对第 1 项值 1 报项, 即 1 个 1, 即 11
 * 第 3 项, 针对第 2 项值 11 报项, 即 2 个 1, 即 21
 * 第 4 项, 针对第 3 项值 21 报项, 即 1 个 2, 1 个 1, 即 1211
 * 第 5 项, 针对第 4 项值 1211 报项, 即 1 个 1, 1 个 2, 2 个 1, 即 111221
 * ...
 * 关键就是提取字符串有多少个连续的字符, 同时纪录该字符和数量
 */

func countAndSay(n int) string {
	if n <= 0 {
		return "0"
	}
	// 初始第 1 项值
	value := "1"
	// n = 3, 只需要循环计算 2 轮
	for num := 1; num < n; num++ {
		// 纪录连续出现的字符以及次数
		countList, strList := []int{}, []string{}
		// 默认出现次数
		count := 1
		for i := 0; i < len(value); i++ {
			if i == 0 {
				// 如果是第一个元素, 放进字符数组
				strList = append(strList, value[i:i+1])
			} else {
				if value[i:i+1] == value[i-1:i] {
					// 如果和前一个元素值相同, 则连续次数 +1
					count++
				} else {
					// 如果不相同, 先将前一个字符的出现次数写入次数数组后, 再将当前字符写入次符数组
					countList = append(countList, count)
					strList = append(strList, value[i:i+1])
					// 重置当前次符出现次数为 1
					count = 1
				}
			}
			if i == len(value)-1 {
				// 如果是最后一个元素, 写入次数数组
				countList = append(countList, count)
			}
		}
		// 拼接次数和字符, 得到新的 value, 代入新一轮计算中
		value = ""
		for i := 0; i < len(countList); i++ {
			value += strconv.FormatInt(int64(countList[i]), 10) + strList[i]
		}
	}
	return value
}
