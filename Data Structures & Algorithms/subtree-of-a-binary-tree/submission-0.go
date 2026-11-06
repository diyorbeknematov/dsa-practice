/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot != nil {
		return false
	}

    if isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) ||
		isSubtree(root.Right, subRoot)
}

func isSameTree(p, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && 
		isSameTree(p.Right, q.Right)
}
