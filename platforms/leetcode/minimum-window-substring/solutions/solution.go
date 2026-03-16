func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	tcount := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tcount[t[i]]++
	}

	required := len(tcount)
	formed := 0

	window := map[byte]int{}

	l := 0
	minLen := len(s) + 1
	start := 0

	for r := 0; r < len(s); r++ {
		c := s[r]
		window[c]++

		if window[c] == tcount[c] {
			formed++
		}

		for l <= r && formed == required {
			if r-l+1 < minLen {
				minLen = r - l + 1
				start = l
			}

			window[s[l]]--
			if window[s[l]] < tcount[s[l]] {
				formed--
			}

			l++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}