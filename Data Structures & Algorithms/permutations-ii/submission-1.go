func permuteUnique(nums []int) [][]int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	var res [][]int
	var backtrack func([]int)	
	backtrack = func(perm []int) {
		if len(perm) == len(nums) {
			res = append(res, append([]int{}, perm...))
			return
		} 

		for v := range freq {
			if freq[v] > 0 {
				perm = append(perm, v)
				freq[v] --

				backtrack(perm)

				freq[v] ++
				perm = perm[:len(perm)-1]
			}
		}
	}

	backtrack([]int{})

	return res
}

