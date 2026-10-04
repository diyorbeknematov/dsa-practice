func mergeAlternately(word1 string, word2 string) string {
	var ans string = ""
	l1 := len(word1)
	l2 := len(word2)
	i, j := 0, 0 

	for i < l1 && j < l2 {
		ans += string(word1[i]) + string(word2[j])
		i++
		j++
	}
	
	for i < l1 {
		ans += string(word1[i])
		i++
	}

	for j < l2 {
		ans += string(word2[j])
		j++
	}

	return ans
}
