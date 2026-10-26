/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int ) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[:n-1]
	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeap{}

	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	dummy := &ListNode{}
	curr := dummy
	for h.Len() != 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node 
		curr = curr.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}


	return dummy.Next
}

