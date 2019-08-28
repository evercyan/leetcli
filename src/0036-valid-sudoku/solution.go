package solution

func isValidSudoku(board [][]byte) bool {
	blen := len(board)
	// 行循环
	for i := 0; i < blen; i++ {
		// 列循环
		for j := 0; j < blen; j++ {
			if string(board[i][j]) == "." {
				continue
			}
			// 检查行重复
			for k := j + 1; k < blen; k++ {
				if board[i][j] == board[i][k] {
					return false
				}
			}
			// 检查列重复
			for k := i + 1; k < blen; k++ {
				if board[i][j] == board[k][j] {
					return false
				}
			}
			// 检查九宫格重复, 当前节点为 (0, 0)
			// 因为校验过行重复, 故 k = 0 这一行不用重复检查, 故 k = i+1, 截止条件到下一次 k % 3 == 0
			// 列的起始 h = j/3*3 = j/3 取整 * 再乘 3
			// 即 当前节点 (0, 0) 时, h = 0, 当 j = 1 的时候, h = 0 -> 截止条件为  j/3*3+3(3 列)
			for k := i + 1; k%3 != 0; k++ {
				for h := j / 3 * 3; h < j/3*3+3; h++ {
					if board[i][j] == board[k][h] {
						return false
					}
				}
			}
		}
	}
	return true
}
