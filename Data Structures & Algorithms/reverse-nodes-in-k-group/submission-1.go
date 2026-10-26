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

	if length < k {
		return head
	}

	dummy := &ListNode{Next: head}
	prevGroup := dummy 
	curr = head

	for length >= k {
		cnt := 1
		tail := curr
		var prev *ListNode 

		for cnt <= k && curr != nil {
			next := curr.Next 
			curr.Next = prev
			prev = curr 
			curr = next 
			cnt++
		}

		prevGroup.Next = prev
		prevGroup = tail

		tail.Next = curr
		length -= k 
	}
	

	return dummy.Next
}


