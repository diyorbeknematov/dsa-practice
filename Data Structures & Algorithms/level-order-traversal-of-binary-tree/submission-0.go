/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
    levels := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		nodes := make([]int, 0)
		for i := 0; i < size; i ++ {
			node := queue[0]
			queue = queue[1:]

			nodes = append(nodes, node.Val) 

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
            if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		levels = append(levels, nodes)
	}

	return levels
}
