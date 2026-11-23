func solveNQueens(n int) [][]string {
	var res [][]string
	
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	cols := make(map[int]bool)
	posDiag := make(map[int]bool) // (r + c)
	negDiag := make(map[int]bool) // (r - c)

	var backtrack func(int) 
	backtrack = func(r int) {
		if r == n {
			solution := make([]string, n)

			for i := 0; i < n; i++ {
				solution[i] = strings.Join(board[i], "")
			}

			res = append(res, solution)
			return 
		}

		for c := 0; c < n; c ++ {
			if cols[c] || posDiag[r+c] || negDiag[r-c] {
				continue
			}

			cols[c] = true
			posDiag[r + c] = true
			negDiag[r - c] = true
			board[r][c] = "Q"

			backtrack(r+1)

			cols[c] = false
			posDiag[r + c] = false
			negDiag[r - c] = false
			board[r][c] = "."
		}
	}

	backtrack(0)

	return res
}

