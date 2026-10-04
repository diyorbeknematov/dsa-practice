func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	m--
	n--

	for i >= 0 {
		if m < 0 {
			nums1[i] = nums2[n]
			n--
		} else if n < 0 {
			nums1[i] = nums1[m]
			m--
		} else if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}

		i--
	}
}

// 10 20 20 40 0 0 
// 1 2 
// 10 20 20 40 0 40 
// 10 20 20 40 20 40 
// 10 20 20 20 20 40 
// 10 20 10 20 20 40 
// 10 2 10 20 20 40 
// 1 2 10 20 20 40 
