type Node struct {
    Key, Val int
    Prev, Next *Node
}

type LRUCache struct {
    capacity int
    nodes    map[int]*Node
    oldest   *Node
    newest   *Node
}

func Constructor(capacity int) LRUCache {
    oldest := &Node{}
    newest := &Node{}

    oldest.Next = newest
    newest.Prev = oldest

    return LRUCache{
        capacity: capacity,
        nodes:    make(map[int]*Node),
        oldest:   oldest,
        newest:   newest,
    }
}

func (this *LRUCache) remove(node *Node) {
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
}

func (this *LRUCache) addToEnd(node *Node) {
    prev := this.newest.Prev

    prev.Next = node
    node.Prev = prev

    node.Next = this.newest
    this.newest.Prev = node
}

func (this *LRUCache) moveToEnd(node *Node) {
    this.remove(node)
    this.addToEnd(node)
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.nodes[key]
    if !ok {
        return -1
    }

    this.moveToEnd(node)
    return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	// if key exists
    if node, ok := this.nodes[key]; ok {
        node.Val = value
        this.moveToEnd(node)
        return
    }

    // Capacity == curCap
    if len(this.nodes) == this.capacity {
        lru := this.oldest.Next
        this.remove(lru)
        delete(this.nodes, lru.Key)
    }

    // Add new value
    node := &Node{
        Key: key,
        Val: value,
    }

    this.addToEnd(node)
    this.nodes[key] = node
}

// oldest <> A <> B <> C <> newest  cap 3 
// put D 
// curCap == cap yes -> oldest.next = oldest.next.next
// prev = this.newest.prev -> prev = C 
// C.next = node 
// node.Prev = C
// node.Next = this.newest 
// this.newest.Prev = node 
