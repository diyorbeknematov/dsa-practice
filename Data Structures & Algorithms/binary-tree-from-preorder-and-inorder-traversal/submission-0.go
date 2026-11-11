/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    mp := make(map[int]int)

	for idx, val := range inorder {
		mp[val] = idx
	}

	var dfs func(preStart, inStart, inEnd int) *TreeNode

	dfs = func(preStart, inStart, inEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}

		rootVal := preorder[preStart]
		rootIdx := mp[rootVal]

		node := &TreeNode{
			Val: rootVal,
		}

		leftSize := rootIdx - inStart

		node.Left = dfs(preStart+1, inStart, rootIdx-1)
		node.Right = dfs(preStart+leftSize+1, rootIdx+1, inEnd)

		return node
	}

	return dfs(0, 0, len(inorder)-1)
}
