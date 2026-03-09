import (
	"strings"
)

func myAtoi(s string) int {
	const INT_MAX = 1<<31 - 1
	const INT_MIN = -1 << 31

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	sign := 1
	if s[0] == '-' {
		fmt.Println(s, s[0], string(s[0]))
		sign = -1
	}
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	val := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			d := int(c) - '0'

			if val > (INT_MAX-d)/10 {
				if sign == 1 {
					return INT_MAX
				}
				return INT_MIN
			}

			val = val*10 + d
		} else {
			break
		}
	}
	val = sign * val
	return val
}