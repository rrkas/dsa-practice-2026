func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var backtrack func(s string, open int, close int)
	backtrack = func(s string, open int, close int) {
		if len(s) == 2*n {
			res = append(res, s)
			return
		}

		if open < n {
			backtrack(s+"(", open+1, close)
		}
		if close < open {
			backtrack(s+")", open, close+1)
		}
	}

	backtrack("", 0, 0)

	return res
}