type Node struct {
	Key, Val int
	Prev, Next *Node
}

type LRUCache struct {
    capacity, curCap int
	nodes map[int]*Node
	oldest, newest *Node
}

func Constructor(capacity int) LRUCache {
	oldest:= &Node{}
	newest := &Node{Prev: oldest}
	oldest.Next = newest


    return LRUCache{
		capacity: capacity,
		nodes: make(map[int]*Node),
		oldest: oldest,
		newest: newest,
	}
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.nodes[key]; !ok {
		return -1
	}
	
	node := this.nodes[key]

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	prev := this.newest.Prev 
	prev.Next = node 
	node.Prev = prev

	node.Next = this.newest
	this.newest.Prev = node 

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodes[key]; ok {
		node.Val = value

		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		prev := this.newest.Prev 
		prev.Next = node 
		node.Prev = prev

		node.Next = this.newest
		this.newest.Prev = node 
		return
	}

	if this.curCap == this.capacity {
		k := this.oldest.Next.Key
		this.oldest.Next = this.oldest.Next.Next
		this.oldest.Next.Prev = this.oldest
		delete(this.nodes, k)
		this.curCap--
	}

	node := &Node{
		Key: key,
		Val: value,
	}

	prev := this.newest.Prev 
	prev.Next = node
	node.Prev = prev

	node.Next = this.newest
	this.newest.Prev = node

	this.nodes[key] = node
	this.curCap++
}

// oldest <> A <> B <> C <> newest  cap 3 
// put D 
// curCap == cap yes -> oldest.next = oldest.next.next
// prev = this.newest.prev -> prev = C 
// C.next = node 
// node.Prev = C
// node.Next = this.newest 
// this.newest.Prev = node 
