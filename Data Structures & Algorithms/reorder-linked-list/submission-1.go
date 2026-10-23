/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    slow := head 
	fast := head 

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next 
		fast = fast.Next.Next 
	}

	curr := slow.Next 
	slow.Next = nil 
	var prev *ListNode

	for curr != nil {
		next := curr.Next 
		curr.Next = prev
		prev = curr
		curr = next 
	}

	curr1 := head 
	curr2 := prev 

	for curr2 != nil {
		leftNext := curr1.Next
		rightNext := curr2.Next

		curr1.Next = curr2
		curr1 = curr1.Next
		curr1.Next = leftNext

		curr1 = curr1.Next 
		curr2 = rightNext
	} 
}

// [0 1 2 3 4 5 6]
// [0 1 2 3] [4 5 6]
// [0 1 2 3] [6 5 4]
// [0 6 1 5 2 4 3]
