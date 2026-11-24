func wordBreak(s string, wordDict []string) []string {
	var res []string 

	wordSet := make(map[string]bool)
	maxWordLen := 0

	for _, word := range wordDict {
		wordSet[word] = true
		maxWordLen = max(maxWordLen, len(word))
	}

	var backtrack func(int, []string) 
	backtrack = func(index int, words []string) {
		if index == len(s) {
			res = append(res, strings.Join(words, " "))
			return 
		}

		end := min(index+maxWordLen, len(s))

		for i := index+1; i <= end; i ++ {
			if wordSet[s[index:i]] {
				words = append(words, s[index:i])
				
				backtrack(i, words)

				words = words[:len(words)-1]
			}
		}
	}

	backtrack(0, []string{})

	return res
}

// s = 'racecariscar   dict=['racecar', 'race', 'car', 'is']
// race 
