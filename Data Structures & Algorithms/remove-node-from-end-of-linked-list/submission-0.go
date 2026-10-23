/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || (head.Next == nil && n >= 1) {
		return nil
	}

	length := 0
	curr := head 
	for curr != nil {
		curr = curr.Next 
		length++
	}

	if length == n {
		return head.Next 
	}

	curr = head 
	for i := 1; i < length - n; i ++ {
		curr = curr.Next 
	}

	curr.Next = curr.Next.Next 
	return head 
}

// 1 2 n = 1
// 
