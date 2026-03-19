func sameChars(a, b string) bool {
	count := make([]int, 26)
	for i := 0; i < len(a); i++ {
		count[a[i]-'a']++
		count[b[i]-'a']--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

func isScramble(s1 string, s2 string) bool {
	memo := make(map[string]bool)

	var dfs func(a, b string) bool
	dfs = func(a, b string) bool {
		key := a + "#" + b

		if val, ok := memo[key]; ok {
			return val
		}

		if a == b {
			memo[key] = true
			return true
		}

		if !sameChars(a, b) {
			memo[key] = false
			return false
		}

		n := len(a)

		for i := 1; i < n; i++ {
			if dfs(a[:i], b[:i]) && dfs(a[i:], b[i:]) {
				memo[key] = true
				return true
			}

			if dfs(a[:i], b[n-i:]) && dfs(a[i:], b[:n-i]) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}

	return dfs(s1, s2)
}