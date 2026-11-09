/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    minVal, maxVal := math.MinInt, math.MaxInt

	var dfs func(node *TreeNode, minVal, maxVal int) bool 
	
	dfs = func (node *TreeNode, minVal, maxVal int) bool {
		if node == nil {
			return true
		}

		if node.Val <= minVal {
			return false
		}

		if node.Val >= maxVal {
			return false
		}

		minVal = min(minVal, node.Val)
		maxVal = max(maxVal, node.Val)

		return dfs(node.Left, minVal, node.Val) && 
			dfs(node.Right, node.Val, maxVal)
	}

	return dfs(root, minVal, maxVal)
}
