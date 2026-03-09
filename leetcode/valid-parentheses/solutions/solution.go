func isValid(s string) bool {
	left := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	right := map[rune]rune{}
	for k, v := range left {
		right[v] = k
	}

	stack := make([]rune, 0)
	for _, c := range s {
		_, ok := left[c]
		if ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			n := len(stack)
			top := stack[n-1]
			stack = stack[:n-1]
			if right[c] != top {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}