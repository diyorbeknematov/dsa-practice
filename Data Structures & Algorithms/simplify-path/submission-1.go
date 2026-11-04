func simplifyPath(path string) string {
	stack := make([]string, 0)

	str := ""
	path += "/"
	for _, s := range path {
		if s == '/' && len(str) != 0 {
			if str == "." {
				str = ""
				continue 
			}

			if str == ".." && len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

			if str != ".." {
				stack = append(stack, str)
			}
			str = ""
		} 

		if s != '/' {
			str += string(s)
		}
	}
	
	if len(stack) == 0 {
		return "/"
	}

	return "/" + strings.Join(stack, "/")
}
