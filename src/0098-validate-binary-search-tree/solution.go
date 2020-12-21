package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 有效的二叉搜索树中序遍历后, 必为升序

func inorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	return res
}

func isValidBST1(root *TreeNode) bool {
	res := inorderTraversal(root)

	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}

	return true
}
