/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	curr := head
	prevNode := head

	for i := 1; i < left; i++ {
		prevNode = curr
		curr = curr.Next
	}

	var prev *ListNode
	tail := curr

	for i := 1; i <= right-left+1; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	if left == 1 {
		head = prev
	} else {
		prevNode.Next = prev
	}

	tail.Next = curr

	return head
}

