func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	letters := []string{
		"abc", "def", "ghi", "jkl",
		"mno", "pqrs", "tuv", "wxyz",
	}

	var res []string 

	var backtrack func(int, string)
	backtrack = func(idx int, str string) {
		if len(str) == len(digits) {
			res = append(res, str)
			return
		}

		if idx >= len(digits) {
			return
		}
		
		i := int(digits[idx] - '2')

		for j := 0; j < len(letters[i]); j ++ {
			str += string(letters[i][j])
			backtrack(idx+1, str)

			str = str[:len(str)-1]
		}
	}

	backtrack(0, "")

	return res
}

// 354    132
// 
