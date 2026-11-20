func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return len(word) == 0
	}

	var backtrack func(int, int, int) bool 

	backtrack = func(row, col int, index int) bool {
		if index == len(word) {
			return true
		}

		if row < 0 || row >= len(board) ||
			col < 0 || col >= len(board[0]) ||
			board[row][col] == '#' ||
			board[row][col] != word[index] {
			return false
		}

		temp := board[row][col]
		board[row][col] = '#'

		found := backtrack(row+1, col, index+1) ||
			backtrack(row, col+1, index+1) ||
			backtrack(row-1, col, index+1) ||
			backtrack(row, col-1, index+1)

		board[row][col] = temp

		return found
	}

	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[0]); j ++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}
