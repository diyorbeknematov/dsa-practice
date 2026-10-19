func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	l, r := 0, m * n - 1
	for l <= r {
		mid := l + (r - l) / 2
		row := mid / n
		col := mid % n 

		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
