/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)

	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		leftRob, leftNotRob := dfs(node.Left)
		rightRob, rightNotRob := dfs(node.Right)

		return (node.Val + leftNotRob + rightNotRob), 
			(max(leftRob, leftNotRob) + max(rightRob, rightNotRob))
	}

	r, nr := dfs(root)

	return max(r, nr)
}
