package solution

func solveSudoku(board [][]byte) {
	solve(board, 0)
}

/**
 * k 是把 board 转换成一维数组后，元素的索引值
 */
func solve(board [][]byte, k int) bool {
	if k == 81 {
		return true
	}
	// 二维数组中的坐标
	r, c := k/9, k%9
	if board[r][c] != '.' {
		// 如果当前值不为 '.', 表示已有数字, 解决下个 key
		return solve(board, k+1)
	}
	// bi, bj 是 rc 所在块的左上角元素的索引值
	bi, bj := r/3*3, c/3*3

	// 匿名函数, 检查元素是否可放置在 [r][c]
	isValid := func(b byte) bool {
		for n := 0; n < 9; n++ {
			// 行检查, 列检查, 九宫格检查
			if board[r][n] == b || board[n][c] == b || board[bi+n/3][bj+n%3] == b {
				return false
			}
		}
		return true
	}

	for b := byte('1'); b <= '9'; b++ {
		if isValid(b) {
			board[r][c] = b
			// 每填入一个数后, 递规去解决下一个 key, 直到所有都满足
			if solve(board, k+1) {
				return true
			}
		}
	}
	board[r][c] = '.'
	return false
}
