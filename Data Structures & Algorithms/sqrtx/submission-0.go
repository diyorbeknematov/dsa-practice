func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 1, x
	ans := 0

	for l <= r {
		mid := l + (r - l) / 2

		if mid <= x / mid {
			ans = mid 
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}
