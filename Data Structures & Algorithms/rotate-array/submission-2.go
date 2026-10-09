func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 1 2 3 4 5 6 7 
// 7 6 5 4 3 2 1 full reverse 
// 5 6 7 4 3 2 1 first k element reverse
// 5 6 7 1 2 3 4 k and len(nums) reverse


