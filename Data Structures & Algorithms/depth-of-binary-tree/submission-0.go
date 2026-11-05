/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)

	return 1 + max(leftMax, rightMax)
}

// 1 -> 2 -> left =  nil 0   1 + max(0, 0) 1 
// 1 -> 2 -> right = nil 0