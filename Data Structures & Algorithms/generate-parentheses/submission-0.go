func generateParenthesis(n int) []string {
	var res []string 

	var dfs func(op, cls int, str string) 

	dfs = func(op, cls int, str string) {
		if cls > op {
			return 
		}

		if len(str) == 2 * n {
			res = append(res, str)
			return
		}

		if op < n {
			dfs(op + 1, cls, str + "(")
		}

		if cls < n {
			dfs(op, cls + 1, str + ")")
		}
	}

	dfs(0, 0, "")

	return res
}



