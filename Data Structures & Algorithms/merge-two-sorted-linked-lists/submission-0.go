/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var dummy *ListNode = &ListNode{} 
	curr := dummy

	curr1, curr2 := list1, list2
	for curr1 != nil && curr2 != nil {
		if curr1.Val < curr2.Val {
			curr.Next = &ListNode{
				Val: curr1.Val,
			}
			curr1 = curr1.Next
		} else {
			curr.Next = &ListNode{
				Val: curr2.Val,
			}
			curr2 = curr2.Next
		}

		curr = curr.Next
	}

	if curr1 != nil {
		curr.Next = curr1
	} else {
		curr.Next = curr2
	}

	return dummy.Next
}
