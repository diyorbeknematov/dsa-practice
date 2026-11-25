func findKthLargest(nums []int, k int) int {
	n := len(nums)
	k = n - k 

	l, r := 0, n 
	for r - l > 1 {
		pivot := nums[l + rand.Intn(r-l)]
		eqL, eqR := partition(nums, l, r, pivot)

		if k < eqL {
			r = eqL
		} else if k >= eqR {
			l = eqR 
		} else {
			l, r = eqL, eqR
			break 
		}
	}

	return nums[l]
}

func partition(arr []int, l, r, pivot int) (int, int) {
	lt, i, gt := l,  l, r 
	for i < gt {
		if arr[i] < pivot {
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
		} else if arr[i] > pivot {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else {
			i++
		}
	}

	return lt, gt
}
