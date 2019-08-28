package solution

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
func maxScoreSightseeingPair(A []int) int {
	if len(A) < 2 {
		return 0
	}
	// cursor, 表示前置的最大的起点 A[i] +i
	score, cursor := 0, 0
	for i := 0; i < len(A); i++ {
		// 当 i 作为终点时, 取 A[j] - j
		score = max(cursor+A[i]-i, score)
		// 当 i 作为起点时, 取 A[i] + i
		cursor = max(cursor, A[i]+i)
	}
	return score
}
