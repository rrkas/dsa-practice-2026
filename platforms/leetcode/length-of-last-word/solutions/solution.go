func lengthOfLastWord(s string) int {
	n := len(s)

	wend := 0
	curr := n - 1
	for curr >= 0 {
		if s[curr] == ' ' {
			if wend < curr {
				curr--
				continue
			} else {
				break
			}
		} else {
			if wend < curr {
				wend = curr
			}
		}
		curr--
	}

	return wend - curr
}