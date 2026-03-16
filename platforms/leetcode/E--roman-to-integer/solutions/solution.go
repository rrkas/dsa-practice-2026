func romanToInt(s string) int {
	cmap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	val := 0
	for i, c := range s {
		if i+1 < len(s) && cmap[c] < cmap[rune(s[i+1])] {
			val -= cmap[c]
		} else {
			val += cmap[c]
		}
	}
	return val
}