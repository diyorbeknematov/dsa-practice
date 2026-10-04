func mergeAlternately(word1 string, word2 string) string {
	ans := make([]byte, 0, len(word1)+len(word2))

	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		ans = append(ans, word1[i], word2[j])
		i++
		j++
	}

	for i < len(word1) {
		ans = append(ans, word1[i])
		i++
	}

	for j < len(word2) {
		ans = append(ans, word2[j])
		j++
	}

	return string(ans)
}
