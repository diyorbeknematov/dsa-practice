/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    var dfs func(node *TreeNode) 
	var ans, n int

	dfs = func(node *TreeNode) {
		if node == nil {
			return 
		}

		if node.Left != nil {
			dfs(node.Left)
		}

		n += 1

		if k == n {
			ans = node.Val 
			return 
		}

		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)

	return ans 
}

//     4
//   3   5
// 2
// 
// 
