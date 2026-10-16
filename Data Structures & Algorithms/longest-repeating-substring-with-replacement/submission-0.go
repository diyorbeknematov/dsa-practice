func characterReplacement(s string, k int) int {
	ans := 0

	freq := make(map[byte]int)
	maxFreq := 0
	
	l := 0
	for r := 0; r < len(s); r++ {

		freq[s[r]]++
		maxFreq = max(maxFreq, freq[s[r]])

		if (r-l+1) - maxFreq > k {
			freq[s[l]]--
			l++
		}
		
		ans = max(ans, r-l+1)

		fmt.Println(maxFreq, string(s[r]))
	}

	return ans
}
