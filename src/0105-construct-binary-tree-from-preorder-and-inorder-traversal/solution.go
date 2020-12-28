package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历 dlr 根左右, preorder = [ 根, [左树前序], [右树前序] ]
// 中序遍历 ldr 左根右, inorder = [ [左树中序], 根, [右树中序] ]
// 从 根 的索引可推导左右树的列表, 递规可得

func search(nums []int, n int) int {
	for key, num := range nums {
		if num == n {
			return key
		}
	}
	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}

	// 前序遍历中的第一个元素即为根
	ret := &TreeNode{
		Val: preorder[0],
	}

	// 根在中序遍历列表中的索引位置, 其左为左树, 其右为右树
	// 注意: 你可以假设树中没有重复的元素, 如果没有此前提, 此做法不可成立
	dIndex := search(inorder, ret.Val)
	if dIndex == -1 {
		return ret
	}

	// 有左子树
	if dIndex > 0 {
		// 左树中序遍历
		lInorder := inorder[:dIndex]
		// 左树前序遍历
		lPreorder := preorder[1 : len(lInorder)+1]
		ret.Left = buildTree(lPreorder, lInorder)
	}

	// 有右子树
	if dIndex < len(inorder)-1 {
		// 右树中序遍历
		rInorder := inorder[dIndex+1:]
		// 右树前序遍历
		rPreorder := preorder[dIndex+1:]
		ret.Right = buildTree(rPreorder, rInorder)
	}

	return ret
}
