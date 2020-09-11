package solution

func numEnclaves(A [][]int) int {
	// 上下左右四个方向 key 偏移
	direction := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	// 纪录节点值为 0, 2 的数量
	count0, count2 := 0, 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 0 {
				count0++
				continue
			}
			// 如果是边界节点, 且节点值为 1
			if (i == 0 || j == 0 || i == len(A)-1 || j == len(A[i])-1) && A[i][j] == 1 {
				// 重置节点值为 2
				A[i][j] = 2
				count2++

				// 写入待处理节点列表
				nodes := [][]int{
					{i, j},
				}
				for {
					if len(nodes) <= 0 {
						break
					}
					// pop 第一个节点
					x, y := nodes[0][0], nodes[0][1]
					if len(nodes) > 1 {
						nodes = append([][]int{}, nodes[1:]...)
					} else {
						nodes = [][]int{}
					}
					// 遍历四个方向, 如果有符合条件的节点, 重置为 2, count2++
					// 再将该节点扔进待处理节点列表
					for _, d := range direction {
						newX, newY := x+d[0], y+d[1]
						if newX >= 0 && newX < len(A) && newY >= 0 && newY < len(A[i]) && A[newX][newY] == 1 {
							A[newX][newY] = 2
							count2++
							nodes = append(nodes, []int{newX, newY})
						}
					}
				}
			}
		}
	}
	return len(A)*len(A[0]) - count0 - count2
}
