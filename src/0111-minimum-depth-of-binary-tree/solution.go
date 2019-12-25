package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	// 如果某一节点为空, 则返回另一节点深度 + 1 即可
	if left == 0 || right == 0 {
		return left + right + 1
	}

	min := 0
	if left > right {
		min = right
	} else {
		min = left
	}
	return min + 1
}
