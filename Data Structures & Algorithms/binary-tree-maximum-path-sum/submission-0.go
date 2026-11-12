/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    maxSum := root.Val 

	var dfs func(node *TreeNode) int 

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := max(0, dfs(node.Left))
		rightMax := max(0, dfs(node.Right))

		maxSum = max(maxSum, leftMax + rightMax + node.Val)

		return node.Val + max(leftMax, rightMax)
	}

	dfs(root)

	return maxSum 
}


//    -15
//   /   \
// 10     20
//        / \
//      15   5
//     /
//   -5
//   / \
//  3  10

//   