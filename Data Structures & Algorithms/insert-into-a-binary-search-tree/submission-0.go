/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    var dfs func(node *TreeNode, val int) *TreeNode

	dfs = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return &TreeNode{Val: val}
		}

		if node.Val > val {
			node.Left = dfs(node.Left, val)
		} else {
			node.Right = dfs(node.Right, val)
		}

		return node
	}
	
	return dfs(root, val)
}
