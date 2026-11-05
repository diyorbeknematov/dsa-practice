/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curr := queue[0]
        queue = queue[1:]

        curr.Left, curr.Right = curr.Right, curr.Left

        if curr.Left != nil {
            queue = append(queue, curr.Left)
        }

        if curr.Right != nil {
            queue = append(queue, curr.Right)
        }
	}

	return root
}
