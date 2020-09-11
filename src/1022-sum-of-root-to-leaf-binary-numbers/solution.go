package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return getSum(root, 0)
}

func getSum(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	// 每递进一层, sum*2 + 节点值
	sum = sum<<1 + node.Val
	// 如果左右子节点均没有, 直接 return
	if node.Left == nil && node.Right == nil {
		return sum
	}
	return getSum(node.Left, sum) + getSum(node.Right, sum)
}
