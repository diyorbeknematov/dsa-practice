func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)

	sort.Ints(candidates)

	var dfs func(int, int)

	dfs = func(i int, total int) {
		if total == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}

		if i >= len(candidates) ||  total > target {
			return 
		}

		cur = append(cur, candidates[i])
		dfs(i+1, total + candidates[i])
		
		j := i + 1
		for ; j < len(candidates); j ++ {
			if candidates[i] != candidates[j] {
				break
			}
		}
		cur = cur[:len(cur)-1]
		dfs(j, total)
	}

	dfs(0, 0)

	return res
}
