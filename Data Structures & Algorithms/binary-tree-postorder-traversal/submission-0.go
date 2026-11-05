/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		dfs(node.Right)

		res = append(res, node.Val)
	}

	dfs(root)

	return res
}
