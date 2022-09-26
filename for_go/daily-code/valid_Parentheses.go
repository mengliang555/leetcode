package dailycode

func isValid(s string) bool {
	if len(s) <= 0 {
		return true
	}
	stack := make([]int32, 0, len(s))
	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) <= 0 || (v == ']' && stack[len(stack)-1] != '[') ||
				(v == ')' && stack[len(stack)-1] != '(') ||
				(v == '}' && stack[len(stack)-1] != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
