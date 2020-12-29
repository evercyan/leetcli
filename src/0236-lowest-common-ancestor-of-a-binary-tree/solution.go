package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 有节点 x 为 p q 最近公共祖先, f(x) 表示 x 包含 p 或 q
//
// 则节点 x 满足以下其一:
// 1. 左子树包含 p 或 q && 右子树包含 p 或 q
// f(x.Left) && f(x.Right)
// 2. x 就是 p 或 q, 并且 x 的左右子树包含 p 或 q
// ( x == p || x == q) && ( f(x.Left) || f(x.Right) )

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// root 必包含 p 和 q, 如果等于其一, 最近祖先即为自己
	if root == p || root == q {
		return root
	}

	// 递规查找左右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 预设 left 和 right 都包含 p, q
	// 如果都不为 nil, 说明都只找到其一, 即 root 为最近祖先
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

/*
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == p.Val || root.Val == q.Val {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {
        return root
    }
    if left == nil {
        return right
    }
    return left
}
*/
