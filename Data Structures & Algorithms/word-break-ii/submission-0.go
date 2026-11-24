type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd bool 
}

func (tn *TrieNode) insert(word string) {
    node := tn

    for _, ch := range word {
        if _, ok := node.Children[ch]; !ok {
            node.Children[ch] = &TrieNode{
                Children: make(map[rune]*TrieNode),
            }
        }

        node = node.Children[ch]
    }

    node.IsEnd = true
}

func (tn *TrieNode) search(word string) bool {
	node := tn 

	for _, ch := range word {
		if _, ok := node.Children[ch]; !ok {
			return false 
		}

		node = node.Children[ch]
	}

	return node.IsEnd
}

func wordBreak(s string, wordDict []string) []string {
	var res []string 

	root := &TrieNode{
		Children: make(map[rune]*TrieNode),
	}

	for _, word := range wordDict {
		root.insert(word)
	}

	var backtrack func(int, []string) 
	backtrack = func(index int, words []string) {
		if index == len(s) {
			res = append(res, strings.Join(words, " "))
			return 
		}

		for i := index; i < len(s); i ++ {
			if root.search(s[index:i+1]) {
				words = append(words, s[index:i+1])
				
				backtrack(i+1, words)

				words = words[:len(words)-1]
			}
		}
	}

	backtrack(0, []string{})

	return res
}

// s = 'racecariscar   dict=['racecar', 'race', 'car', 'is']
// race 
