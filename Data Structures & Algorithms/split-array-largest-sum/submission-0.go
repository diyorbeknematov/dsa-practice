func splitArray(nums []int, k int) int {
	lo, hi := 0, 0
	for _, v := range nums {
		lo = max(lo, v)
		hi += v
	}

	for lo < hi {
		mid := lo + (hi-lo)/2
		if fitsInParts(nums, k, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func fitsInParts(nums []int, k, maxSum int) bool {
	parts, sum := 1, 0
	for _, v := range nums {
		if sum+v > maxSum {
			parts++
			if parts > k {
				return false
			}
			sum = v
		} else {
			sum += v
		}
	}
	return true
}
