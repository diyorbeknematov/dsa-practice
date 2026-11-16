type TrieNode struct {
	Val rune
	Children [26]*TrieNode
	Index int 
}

type Trie struct {
	root *TrieNode
}

func (this *Trie) add(word string, index int) {
    node := this.root

    for _, w := range word {
        idx := w - 'a'

        if node.Children[idx] == nil {
            node.Children[idx] = &TrieNode{
                Val:   w,
                Index: -1,
            }
        }

        node = node.Children[idx]
    }

    node.Index = index
}

func findWords(board [][]byte, words []string) []string {
	prefTree := &Trie{
		root: &TrieNode{Index: -1},
	}

	for i, word := range words {
		prefTree.add(word, i)
	}

	res := []string{}
	var dfs func(row, col int, node *TrieNode)

	dfs = func(row, col int, node *TrieNode) {
		
		if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
			return
		}

		if board[row][col] == '#' {
			return
		}

		idx := board[row][col] - 'a'
		next := node.Children[idx]

		if next == nil {
			return
		}

		if next.Index != -1 {
			res = append(res, words[next.Index])
			next.Index = -1
		}

		ch := board[row][col]
		board[row][col] = '#'

		dfs(row + 1, col, next)
		dfs(row, col + 1, next)
		dfs(row, col - 1, next)
		dfs(row - 1, col, next)

		board[row][col] = ch
	}

	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[0]); j ++ {
			dfs(i, j, prefTree.root)
		}
	}

	return res
}
