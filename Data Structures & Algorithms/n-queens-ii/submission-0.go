func totalNQueens(n int) int {
	cols := make([]bool, n)
	posDiag := make([]bool, 2*n - 1) // (r - c + n - 1)
	negDiag := make([]bool, 2*n - 1) // (r + c)

	var backtrack func(row int) int
	backtrack = func(row int) int {
		if row == n {
			return 1
		}

		cnt := 0

		for col := 0; col < n; col++ {
			if cols[col] || posDiag[row-col+n-1] || negDiag[row+col] {
				continue
			}

			cols[col] = true
			posDiag[row-col+n-1] = true
			negDiag[row+col] = true

			cnt += backtrack(row + 1)

			cols[col] = false
			posDiag[row-col+n-1] = false
			negDiag[row+col] = false
		}

		return cnt
	}

	return backtrack(0)
}
