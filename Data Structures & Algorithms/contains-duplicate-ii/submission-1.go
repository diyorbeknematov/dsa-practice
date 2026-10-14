func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, val := range nums {
		if _, ok := seen[val]; ok {
			if abs(i-seen[val]) <= k {
				return true
			}
		}

		seen[val] = i
	}

	return false
}

func abs(num int) int {
	
	if num < 0 {
		return num * (-1)
	}

	return num
}
