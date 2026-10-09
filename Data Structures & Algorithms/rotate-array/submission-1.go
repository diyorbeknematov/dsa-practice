func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	arr := make([]int, n)
	indx := 0
	for i:=n-k; i < n; i++ {
		arr[indx] = nums[i]
		indx++
	}
	
	for i := 0; i < n - k; i++ {
		arr[indx] = nums[i]
		indx++
	}

	for i, num := range arr {
		nums[i] = num
	}
}

// 1 2 3 4 5 6 7 8  k=4
// 8 1 2 3 4 5 6 7  k=3
// 7 8 1 2 3 4 5 6  k=2
// 6 7 8 1 2 3 4 5  k=1
// 5 6 7 8 1 2 3 4  k=0 

// 