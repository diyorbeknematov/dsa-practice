func wordBreak(s string, wordDict []string) []string {
	var res []string 

	memo := make(map[int][][]string)
	wordSet := make(map[string]bool)

	for _, word := range wordDict {
		wordSet[word] = true
	}

	var backtrack func(int) [][]string
	backtrack = func(index int) [][]string {
		if index == len(s) {
			return [][]string{{}}
		}

		if cached, ok := memo[index]; ok {
			return cached
		} 

		result := make([][]string, 0)

		for i := index + 1; i <= len(s); i++ {
			word := s[index:i]

			if wordSet[word] {
				comb := backtrack(i)

				for _, c := range comb {
					newComb := append([]string{word}, c...)
					result = append(result, newComb)
				}
			}
		}

		memo[index] = result
    	return result
	}

	words := backtrack(0)
	for _, w := range words {
		res = append(res, strings.Join(w, " "))
	}

	return res
}

// s = 'racecariscar   dict=['racecar', 'race', 'car', 'is']
// race 
