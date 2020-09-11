package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 后序遍历, 即 lrd
 */
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 出现了新的层级, 后序遍历, 插至切片头部
		if level >= len(res) {
			res = append([][]int{{}}, res...)
		}
		// 写入元素值
		// 如果 len(res) = 4, 则 level = 1 表示第二层, 写入 key = 2
		levelKey := len(res) - level - 1
		res[levelKey] = append(res[levelKey], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return res
}
