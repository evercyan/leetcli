package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	num := sum - root.Val
	// 已经到了叶子节点了
	if root.Left == nil && root.Right == nil {
		return num == 0
	}
	return hasPathSum(root.Left, num) || hasPathSum(root.Right, num)
}
