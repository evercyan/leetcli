package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs 遍历数据按 level 存储即可

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// 需要处理一下 res[level] 的初始化
		if len(res) < level+1 {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	return res
}
