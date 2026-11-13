type PrefixTree struct {
	Children map[rune]*PrefixTree
	IsEnd bool
}

func Constructor() PrefixTree {
    return PrefixTree{
        Children: make(map[rune]*PrefixTree),
    }
}

func (this *PrefixTree) Insert(word string) {
    node := this

    for _, w := range word {
        if _, ok := node.Children[w]; !ok {
            node.Children[w] = &PrefixTree{
                Children: make(map[rune]*PrefixTree),
            }
        }

        node = node.Children[w]
    }

    node.IsEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	node := this 

	for _, w := range word {
		if _, ok := node.Children[w]; !ok {
			return false
		}

		node = node.Children[w]
	}

	return node.IsEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this 

	for _, w := range prefix {
		if _, ok := node.Children[w]; !ok {
			return false
		}

		node = node.Children[w]
	}

	return true
}
