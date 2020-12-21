package solution

// n*n 的矩阵, 顺时针转换
// 坐标转换为: (i, j) => (j, n-1-i)

func rotate(matrix [][]int) {
	n := len(matrix)

	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ret[j][n-1-i] = matrix[i][j]
		}
	}

	copy(matrix, ret)
}
