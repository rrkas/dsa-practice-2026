import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minl := len(strs[0])
	for _, s := range strs {
		if len(s) < minl {
			minl = len(s)
		}
	}
	fmt.Println(minl)
	if minl == 0 {
		return ""
	}

	cs := strings.Builder{}

	for i := 0; i < minl; i++ {
		c0 := strs[0][i]
		mm := false
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != c0 {
				mm = true
				break
			}
		}
		if !mm {
			cs.WriteByte(c0)
		} else {
			break
		}
	}

	return cs.String()
}