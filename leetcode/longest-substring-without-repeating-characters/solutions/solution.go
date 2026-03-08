func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]bool)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}

		seen[s[right]] = true

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}