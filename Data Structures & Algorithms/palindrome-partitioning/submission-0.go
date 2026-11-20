func partition(s string) [][]string {
	var res [][]string

	var backtrack func(int, []string) 
	backtrack = func(start int, path []string) {
		// base case
		if start == len(s) {
			copyPath := append([]string{}, path...)
			res = append(res, copyPath)
    		return
		}

		for i := start; i < len(s); i ++ {
			sub := s[start:i+1]

			if isPalindrome(sub) {
				path = append(path, sub)

				backtrack(i+1, path)

				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0, []string{})

	return res
}

func isPalindrome(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}

	return true
}

//                []
//          a          ''
//         aa a       a    ''
//       aab   ab    ab   b   ''