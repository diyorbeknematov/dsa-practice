/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    var build func(grid[][]int, row, col, size int) *Node

	build = func(grid[][]int, row, col, size int) *Node {
		isSameVal := true
		num := grid[row][col]
		for i := row; i < row + size; i ++ {
			for j := col; j < col + size; j ++ {
				if num != grid[i][j] {
					isSameVal = false
					break
				}
			}
			if !isSameVal {
				break
			}
		}

		if !isSameVal {
			node := &Node{
				IsLeaf: false,
			}

			half := size / 2

			node.TopLeft = build(grid, row, col, half)
			node.TopRight = build(grid, row, col + half, half )
			node.BottomLeft = build(grid, row + half, col, half)
			node.BottomRight = build(grid, row + half, col + half, half)

			return node
		}

		val := num == 1

		node := &Node{
			IsLeaf: true,
			Val: val,
		}

		return node
	}

	return build(grid, 0, 0, len(grid))
}
