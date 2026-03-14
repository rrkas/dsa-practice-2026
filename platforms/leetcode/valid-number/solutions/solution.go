func isNumber(s string) bool {
	seenDigit := false
	seenDot := false
	seenExp := false
	digitAfterExp := true

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c >= '0' && c <= '9' {
			seenDigit = true
			if seenExp {
				digitAfterExp = true
			}

		} else if c == '+' || c == '-' {
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}

		} else if c == '.' {
			if seenDot || seenExp {
				return false
			}
			seenDot = true

		} else if c == 'e' || c == 'E' {
			if seenExp || !seenDigit {
				return false
			}
			seenExp = true
			digitAfterExp = false

		} else {
			return false
		}
	}

	return seenDigit && digitAfterExp
}