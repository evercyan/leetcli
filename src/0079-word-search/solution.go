package solution

// 回溯算法

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])

	// 同个单元格不能重复使用, 需记录使用状态
	used := make([][]bool, row)
	for k, _ := range board {
		used[k] = make([]bool, col)
	}

	// k 表示 word 的字符索引, 比对完一个字符后, 继续 k+1
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// 表示 word 已比对完
		if k == len(word) {
			return true
		}

		if i < 0 || j < 0 || i >= row || j >= col {
			return false
		}
		if used[i][j] || board[i][j] != word[k] {
			return false
		}

		// 做出选择
		used[i][j] = true

		// 上下左右继续比对下个字符
		// 此处需要注意四个条件满足其一即可, 故用 ||, 避免耗时
		tmp := dfs(i, j-1, k+1) || dfs(i, j+1, k+1) || dfs(i-1, j, k+1) || dfs(i+1, j, k+1)

		// 撤消选择
		used[i][j] = false

		return tmp
	}

	// 比对找遍历起始点
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
