type Node struct {
    Key, Val, Freq int
    Prev, Next     *Node
}

type DoublyLinkedList struct {
    oldest, newest *Node
}

type LFUCache struct {
    capacity int
    nodes    map[int]*Node
    freqList map[int]*DoublyLinkedList
    minFreq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		nodes: make(map[int]*Node),
		freqList: make(map[int]*DoublyLinkedList),
	}
}

func (this *LFUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LFUCache) addToEnd(nodeList *DoublyLinkedList, node *Node) {
    prev := nodeList.newest.Prev

	prev.Next = node
	node.Prev = prev 

	node.Next = nodeList.newest
	nodeList.newest.Prev = node 
}

func (this *LFUCache) moveToUpFreq(node *Node) {
    this.remove(node)
    node.Freq++

    if this.freqList[node.Freq] == nil {
        oldest := &Node{}
        newest := &Node{Prev: oldest}
        oldest.Next = newest

        this.freqList[node.Freq] = &DoublyLinkedList{
            oldest: oldest,
            newest: newest,
        }
    }

    this.addToEnd(this.freqList[node.Freq], node)
}

func (this *LFUCache) isEmpty (nodeList *DoublyLinkedList) bool {
	return nodeList.oldest.Next == nodeList.newest
}


func (this *LFUCache) Get(key int) int {
    if _, ok := this.nodes[key]; !ok {
		return -1
	}

	node := this.nodes[key]
	freq := node.Freq
	nodeList := this.freqList[node.Freq]

	this.moveToUpFreq(node)
	if this.isEmpty(nodeList) && this.minFreq == freq {
		this.minFreq++
	}

	return node.Val
}


func (this *LFUCache) Put(key int, value int)  {
	if this.capacity == 0 {
		return
	}
	
	if node, ok := this.nodes[key]; ok {
		node.Val = value
		this.moveToUpFreq(node)
		return
	}

	if this.capacity == len(this.nodes) {
		nodeList := this.freqList[this.minFreq]
		k := nodeList.oldest.Next.Key
		this.remove(nodeList.oldest.Next)
		delete(this.nodes, k)
	}

	node := &Node{
		Key: key,
		Val: value,
		Freq: 1,
	}

	this.minFreq = 1
	if this.freqList[this.minFreq] == nil {
		oldest := &Node{}
        newest := &Node{Prev: oldest}
        oldest.Next = newest

        this.freqList[node.Freq] = &DoublyLinkedList{
            oldest: oldest,
            newest: newest,
        }
	}
	this.addToEnd(this.freqList[this.minFreq], node)
	this.nodes[key] = node
}	


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param1 := obj.Get(key);
 * obj.Put(key,value);
 */
