func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]bool)

	i, j := 0, 0 
	for j < len(nums) {
		if abs(i-j) > k {
			set[nums[i]] = false
			i++
		}

		if set[nums[j]] {
			return true
		}

		set[nums[j]] = true
		j++
	}

	return false
}

func abs(num int) int {
	
	if num < 0 {
		return num * (-1)
	}

	return num
}
