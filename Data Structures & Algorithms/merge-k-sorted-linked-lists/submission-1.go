/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil 
	}

	for i := 1; i < len(lists); i *= 2 {
		for j := 0; j + i < len(lists); j += i * 2 {
			lists[j] = merge(lists[j], lists[j+i])
		}
	}
	
	return lists[0]
}

func merge(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy 

	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next

		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return dummy.Next
}
