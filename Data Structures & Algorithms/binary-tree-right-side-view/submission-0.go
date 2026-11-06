/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
    ans := make([]int, 0)
	queue := []*TreeNode{ root }

	for len(queue) > 0 {
		qLen := len(queue)

		ans = append(ans, queue[qLen-1].Val)

		for i := 0; i < qLen; i ++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return ans
}
