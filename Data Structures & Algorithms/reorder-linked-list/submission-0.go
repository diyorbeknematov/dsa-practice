/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    curr := head 
	nodes := make([]*ListNode, 0)

	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next 
	}

	dummy := &ListNode{}
	curr = dummy 

	l, r := 0, len(nodes)-1
	for l <= r {
		if l == r {
			curr.Next = nodes[l]
			curr = curr.Next 
			break
		}

		curr.Next = nodes[l]
		curr = curr.Next 
		curr.Next = nodes[r]
		curr = curr.Next 
		l++
		r--
	}
	curr.Next = nil 
	head = dummy.Next
}
