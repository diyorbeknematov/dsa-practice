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
			Next: curr.Next,
			Random: curr.Random,
		}

		oldToNew[curr] = node
		curr = curr.Next
	}

	curr1 := head 

	dummy := &Node{}
	curr2 := dummy 

	for curr1 != nil {
		node := oldToNew[curr1]
		node.Next = oldToNew[curr1.Next]
		node.Random = oldToNew[curr1.Random]
		curr2.Next = node 

		curr1 = curr1.Next 
		curr2 = curr2.Next
	}

	return dummy.Next 
}




