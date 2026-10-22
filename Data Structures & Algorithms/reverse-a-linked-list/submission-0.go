/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
	curr := head 

	for curr != nil {
		next := curr.Next 
		curr.Next = prev 
		prev = curr 
		curr = next 
	}

	return prev
}

// head 0 - 1 - 2 - 3 
// curr 0 - 1 - 2 - 3 
// p = nil 
// 
// next = curr.next   1 2 3 
// curr.next = p  curr = 0 nil 
// p = curr  p = 0 nil 
// curr = next  curr = 1 2 3 

// next = 2 3 
// curr.next = p = 1 0 nil 
// p = curr = 1 0 nil 
// curr = next = 2 3 