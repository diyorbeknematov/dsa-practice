/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    first := l1 
	second := l2 

	dummy := &ListNode{}
	curr := dummy 

	digit := 0 

	for first != nil || second != nil {
		d := 0
		if first != nil {
			d += first.Val 
			first = first.Next 
		}

		if second != nil {
			d += second.Val 
			second = second.Next
		}

		curr.Next = &ListNode{
			Val: (d + digit) % 10,
		}

		curr = curr.Next 
		digit = (d + digit) / 10
	}

	if digit != 0 {
		curr.Next = &ListNode{Val: digit}
	}

	return dummy.Next 
}
// 6 2 3 
// 5 7 9 
// 

// (5 + 6 + 0) % 10  



