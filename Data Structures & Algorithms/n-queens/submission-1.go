func solveNQueens(n int) [][]string {
	var res [][]string
	
	cols := make([]int, n)
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	var backtrack func(int, [][]string) 
	backtrack = func(row int, board [][]string) {
		if row == n {
			solution := make([]string, n)

			for i := 0; i < n; i++ {
				solution[i] = strings.Join(board[i], "")
			}

			res = append(res, solution)
			return 
		}
		for i := 0; i < n; i ++ {
			board[row][i] = "Q"
			cols[row] = i
			if isPlace(cols, row, i, n) {
				backtrack(row+1, board)
			}

			board[row][i] = "."
		}
	}

	backtrack(0, board)

	return res
}

func isPlace(cols []int, row, col, n int) bool {
	for i := 0; i < row; i++ {
		prevCol := cols[i]

		if prevCol == col || abs(i-row) == abs(prevCol-col) {
			return false
		}
	}

	return true
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}

	return num
}
