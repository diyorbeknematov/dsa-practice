func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

    maxLength := 0
	length := 1
	prev := nums[0]

	for _, num := range nums {
		if num-prev == 1 {
			length++
			prev = num
		} else if num-prev > 1 {
			maxLength = max(maxLength, length)
			length = 1
			prev = num
		}
	}

	return max(maxLength, length)
}

// 2 20 4 10 3 4 5
// 
