func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	row := 0
	l, r := 0, m - 1
	for l <= r {
		mid := l + (r - l) / 2

		if matrix[mid][0] <= target && matrix[mid][n-1] >= target {
			row = mid 
			break
		}

		if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	l, r = 0, n - 1
	for l <= r {
		mid := l + (r - l) / 2

		if matrix[row][mid] == target {
			return true
		}

		if matrix[row][mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
