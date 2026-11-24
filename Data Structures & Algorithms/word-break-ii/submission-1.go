func wordBreak(s string, wordDict []string) []string {
	var res []string 

	wordSet := make(map[string]bool)

	for _, word := range wordDict {
		wordSet[word] = true
	}

	var backtrack func(int, []string) 
	backtrack = func(index int, words []string) {
		if index == len(s) {
			res = append(res, strings.Join(words, " "))
			return 
		}

		for i := index; i < len(s); i ++ {
			if wordSet[s[index:i+1]] {
				words = append(words, s[index:i+1])
				
				backtrack(i+1, words)

				words = words[:len(words)-1]
			}
		}
	}

	backtrack(0, []string{})

	return res
}

// s = 'racecariscar   dict=['racecar', 'race', 'car', 'is']
// race 
