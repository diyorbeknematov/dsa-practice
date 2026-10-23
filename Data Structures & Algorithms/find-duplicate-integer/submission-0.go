func findDuplicate(nums []int) int {
    freq := make([]int, len(nums))

	for _, v := range nums {
		freq[v]++
	}

	for num, v := range freq {
		if v > 1 {
			return num
		}
	}

	return 0
}
