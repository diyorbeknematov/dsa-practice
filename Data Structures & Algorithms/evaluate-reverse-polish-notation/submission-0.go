func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens { 
		switch token {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			stack = append(stack, operate(a, b, token))
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func operate(num1, num2 int, op string) int {
	var ans int

	switch op {
	case "+": ans = num1 + num2
	case "-": ans = num1 - num2
	case "*": ans = num1 * num2
	case "/": ans = num1 / num2
	}

	return ans
}
