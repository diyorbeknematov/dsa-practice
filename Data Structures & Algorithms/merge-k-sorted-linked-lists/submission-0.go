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

	if len(lists) == 1 {
		return lists[0]
	}

	l, r := 0, len(lists)-1
	nodes := make([]*ListNode, 0, len(lists)/2)

	for l < r {
		curr1 := lists[l]
		curr2 := lists[r]

		dummy := &ListNode{}
		curr := dummy

		for curr1 != nil && curr2 != nil {
			if curr1.Val < curr2.Val {
				curr.Next = curr1
				curr = curr.Next

				curr1 = curr1.Next
			} else {
				curr.Next = curr2
				curr = curr.Next 

				curr2 = curr2.Next
			}
		}

		if curr1 != nil {
			curr.Next = curr1
		} else {
			curr.Next = curr2
		}

		nodes = append(nodes, dummy.Next)
		l++
		r--
	}

	if l == r {
		nodes = append(nodes, lists[l])
	}

	return mergeKLists(nodes)
}
