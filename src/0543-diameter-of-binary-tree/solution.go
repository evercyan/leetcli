package solution

// 经过根节点, 即为: lDeep + rReep
// 不经过根节点, 即为: 计算第一个节点的 (lDeep + rDeep + 1), 取 max

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ret int

	var ldr func(root *TreeNode) int
	ldr = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 左右子树的最大节点数
		l, r := ldr(root.Left), ldr(root.Right)
		// 直径比对
		ret = max(ret, l+r+1)
		return max(l, r) + 1
	}
	ldr(root)

	if ret == 0 {
		return 0
	}
	return ret - 1
}
