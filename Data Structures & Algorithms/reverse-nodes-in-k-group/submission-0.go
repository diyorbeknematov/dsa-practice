/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    length := 0
	curr := head 
	for curr != nil {
		length++
		curr = curr.Next 
	}

	return reverseList(head, k, length)
}

func reverseList(head *ListNode, k, length int) *ListNode {
	if length < k {
		return head
	}

	dummy := &ListNode{Next: head}
 
	cnt := 1
	curr := head
	tail := head 

	var prev *ListNode 

	for cnt <= k && curr != nil {
		next := curr.Next 
		curr.Next = prev
		prev = curr 
		curr = next 
		cnt++
	}

	dummy.Next = prev 
	tail.Next = reverseList(curr, k, length-k)

	return dummy.Next
}

