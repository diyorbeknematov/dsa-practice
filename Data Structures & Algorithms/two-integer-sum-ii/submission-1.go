func twoSum(numbers []int, target int) []int {
	n := len(numbers)

	for i := 0; i < n; i ++ {
		index2 := binarySearch(numbers, i+1, n-1, target-numbers[i])
		if index2 != -1 {
			return []int{i+1, index2+1}
		}
	}

	return []int{}
}

func binarySearch(nums []int, l, r, target int) int {
	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

