/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    var dfs func(node *TreeNode, maxVal int)

	cnt := 0
	dfs = func(node *TreeNode, maxVal int) {
		if node == nil {
			return 
		}

		if maxVal <= node.Val {
			cnt ++
		} 

		maxVal = max(maxVal, node.Val)

		if node.Left != nil {
			dfs(node.Left, maxVal)
		}

		if node.Right != nil {
			dfs(node.Right, maxVal)
		}
	}

	dfs(root, root.Val)
	return cnt
}
