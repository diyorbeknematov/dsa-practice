func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	cur := make([]int, 0)

	var dfs func(int, []int)

	dfs = func(start int, cur []int) {
		if len(cur) == k {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return 
		}

		if start > n {
			return 
		}

		cur = append(cur, start)
		dfs(start+1, cur)

		cur = cur[:len(cur)-1]
		dfs(start+1, cur)
	}

	dfs(1, cur)
	return res
}

// 1 2 3   k = 3
// i       
//       []
//     [1]                            []
// [1,2] [1]                       [2]                     []
//[1,2,3] [1,2]                  [2,3] [2]              [3]  []
// r   [1,2,4]  [1,2]         [2,3,4] [2,3]         [3,4] [3]
//      r   [1,2,5]  [1,2]      r   [2,3,5] [2,3]  [3,4,5] [3,4]
//            r        r               r      r       r      r 
