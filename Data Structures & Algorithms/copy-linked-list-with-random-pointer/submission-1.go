/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToNew := make(map[*Node]*Node)

	curr := head 
	for curr != nil {
		node := &Node{
			Val: curr.Val,
		}

		oldToNew[curr] = node
		curr = curr.Next
	}

	curr1 := head 
	curr2 := oldToNew[head] 

	for curr2 != nil {
		curr2.Next = oldToNew[curr1.Next]
		curr2.Random = oldToNew[curr1.Random]

		curr2 = curr2.Next 
		curr1 = curr1.Next
	}

	return oldToNew[head]
}




