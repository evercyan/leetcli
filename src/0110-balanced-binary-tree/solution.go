package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	max := 0
	if left > right {
		max = left
	} else {
		max = right
	}
	return max + 1
}

/**
 * 自下而上, 每个节点只计算一次 O(n)
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isBalanced(root.Left) && isBalanced(root.Right) {
		num := maxDepth(root.Left) - maxDepth(root.Right)
		if num > 1 || num < -1 {
			return false
		}
		return true
	}
	return false
}

/**
 * 自上而下, 不断递归左右子树
 * O(n^2)
 */
// func isBalanced1(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	num := maxDepth(root.Left) - maxDepth(root.Right)
// 	if num > 1 || num < -1 {
// 		return false
// 	}
// 	return isBalanced(root.Left) && isBalanced(root.Right)
// }
