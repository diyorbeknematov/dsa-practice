func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r - l) / 2
		
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

// -1 0 2 4 6 8  l=0   r=5    mid=2  
// -1 0 2 4 6 8  l4