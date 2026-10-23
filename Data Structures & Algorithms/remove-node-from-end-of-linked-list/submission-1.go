/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

    dummy := &ListNode{
		Next: head,
	}

	fast, slow := dummy, dummy
	for i := 0; i < n; i ++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next 
		slow = slow.Next 
	}

	slow.Next = slow.Next.Next 
	return dummy.Next 
}

// 1 -> 2 -> 3 -> 4 -> 5    n=2
// 0 -> 1 -> 2 -> 3 -> 4 -> 5
// fast, slow = dummy, dummy
// fast n step forward  fast = 2 -> 3 -> 4 -> 5
// slow = 0 -> 1 -> 2 -> 3 -> 4 -> 5 

// for fast.Next != nil 
// fast 3->4->5       | 4->5       | 5       | 5->nil break 
// slow 1->2->3->4->5 | 2->3->4->5 | 3->4->5 | 

// slow.next = slow.next.next 3.next = 3.next.next = 3->5 


// 1->2->3     n=3
// fast = dummy = 0->1->2->3   3 steps 3 
// slow.next = slow.next.next