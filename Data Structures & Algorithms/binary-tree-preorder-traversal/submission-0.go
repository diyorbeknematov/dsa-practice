/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
    res := make([]int, 0)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		res = append(res, curr.Val)

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}

	return res
}
