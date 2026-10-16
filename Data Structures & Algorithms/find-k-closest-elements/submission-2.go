func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1

	for (r - l + 1) > k {
		if abs(arr[l]-x) > abs(arr[r]-x) {
			l++
		} else if abs(arr[l]-x) == abs(arr[r]-x) && arr[l] <= arr[r] {
			r--
		} else {
			r--
		}
	}

	ans := make([]int, 0, k)
	
	for l <= r {
		ans = append(ans, arr[l])
		l++
	}

	return ans
}

func abs(num int) int {
	if num < 0 {
		return (-1)*num
	}

	return num
}
