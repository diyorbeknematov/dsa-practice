/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    var is bool = true
	var dfs func(root *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		
		if abs(left - right) > 1 {
			is = false
		}

		return 1 + max(left, right)
	}

	dfs(root)
	return is
}

func abs(num int) int {
	if num < 0 {
		return (-1) * num
	}

	return num
}
