func decodeString(s string) string {
	currentString := ""
	countStack := make([]int, 0)
	stringStack := make([]string, 0)

	i := 0
	for  i < len(s) {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			k := 0
			for i < len(s) && 48 <= s[i] && s[i] <= 57 {
				k = k * 10 + int(s[i]-'0')
				i++
			}

			if k != 0 {
				countStack = append(countStack, k)
			}
			continue

		case s[i] == '[':
			stringStack = append(stringStack, currentString)
			currentString = ""

		case s[i] == ']':
			k := countStack[len(countStack)-1]
			str := ""
			for j := 0; j < k; j ++ {
				str += currentString
			}

			currentString = stringStack[len(stringStack)-1] + str

			countStack = countStack[:len(countStack)-1]
			stringStack = stringStack[:len(stringStack)-1]

		default:
			currentString += string(s[i])
		}

		i++
	}

	return currentString
}

// axb3[z]4[c]
// axb 3[z]
