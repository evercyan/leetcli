package solution

// 访问过的陆地被置为 0, 并重置其相连路径

func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])

	var count int
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if grid[x][y] == '1' {
				count++
				dfs(grid, x, y)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, x, y int) {
	// 越过边界 || 相连不为陆地
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'

	// 继续遍历当前陆地四周
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
}
