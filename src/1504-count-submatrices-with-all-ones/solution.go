package solution

/**
 * 思路:
 * 任意 [i][j], 以其为左上角, 向右下方向统计其所能形成的非 1*1(不包括自己) 的矩形数量
 */

func numSubmat(mat [][]int) int {
	row, col := len(mat), len(mat[0])
	num := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 如果为 0, 则当前节点无此
			if mat[i][j] == 0 {
				continue
			}
			// colBadge 为矩形右边界
			colBadge := col
			// 以 [p][q] 为矩形右下角
			for p := i; p < row; p++ {
				for q := j; q < col && q < colBadge; q++ {
					if mat[p][q] == 1 {
						num++
					} else {
						colBadge = q
					}
				}
			}
		}
	}
	return num
}
