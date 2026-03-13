func isMatch(s string, p string) bool {
	sn, pn := len(s), len(p)
	i, j := 0, 0
	star, match := -1, 0

	for i < sn {
		if j < pn && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < pn && p[j] == '*' {
			star = j
			match = i
			j++
		} else if star != -1 {
			j = star + 1
			match++
			i = match
		} else {
			return false
		}
	}

	for j < len(p) && p[j] == '*' {
		j++
	}

	return j == pn
}