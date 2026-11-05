/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    stack := []*TreeNode{}
    curr := root
    var prev *TreeNode

    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }

        top := stack[len(stack)-1]

        if top.Right != nil && top.Right != prev {
            curr = top.Right
        } else {
            res = append(res, top.Val)
            prev = top
            stack = stack[:len(stack)-1]
        }
    }

    return res
}
