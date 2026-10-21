func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r - l) / 2

		if nums[mid] == target {
			return true
		}

		if nums[mid] > nums[l] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[l] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			l++
		}
	}

	return false
}

// 3 4 4 5 1 2 2 

// if nums[mid] > nums[l] yes left sorted  
	// if nums[l] =< target && target < nums[mid]: r = mid - 1
	// else: l = mid + 1
// if nums[mid] < nums[l] yes right sorted 
	// if nums[mid] < target && target <= nums[r]: l=mid+1
	// else: r=mid-1
// 
// else: means if nums[mid] == nums[l]: l++
// 
