func minExtraChar(s string, dictionary []string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 0

	for i := 1; i <= len(s); i++ {
		dp[i] = len(s)
	}

	words := make(map[string]bool)

	for _, word := range dictionary {
		words[word] = true
	}

	for i := 0; i < len(s); i ++ {
		dp[i+1] = min(dp[i+1], dp[i]+1)
		
		for j := i; j < len(s); j ++ {
			if words[s[i:j+1]] {
				dp[j+1] = min(dp[j+1], dp[i])
			} 
		}
	}

	return dp[len(s)]
}
