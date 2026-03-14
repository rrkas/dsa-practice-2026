func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)

	if m == 0 {
		return 0
	}

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && haystack[i+j] == needle[j] {
			j++
		}
		if j == m {
			return i
		}
	}

	return -1
}