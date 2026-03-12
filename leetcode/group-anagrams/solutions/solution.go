func groupAnagrams(strs []string) [][]string {
	freqmap := make(map[[26]int][]string)
	for _, word := range strs {
		count := [26]int{}
		for _, c := range word {
			count[c-'a']++
		}
		freqmap[count] = append(freqmap[count], word)
	}

	res := make([][]string, 0, len(freqmap))
	for _, wl := range freqmap {
		res = append(res, wl)
	}
	return res
}