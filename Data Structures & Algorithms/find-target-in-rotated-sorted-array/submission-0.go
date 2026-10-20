func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r - l) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[r]  {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

// 12

// 1 2 3 4 5 6 7 8 9 10 11 12 13 
// 4 5 6 7 8 9 10 11 12 13 1 2 3 

// l=0 r=12 mid=0+(12-0)/2=6  nums[6] > nums[12] yes nums[6] < target yes 
// l=mid+1=7 r=12 mid=7+(12-7)/2=9  nums[9] > nums[12] yes nums[9] > target yes 
// l=7 r=mid-1=8  mid=7+(8-7)/2=7   nums[7] > nums[8] no nums[mid] > target yes 
// l=mid+1=8 r=8  mid


