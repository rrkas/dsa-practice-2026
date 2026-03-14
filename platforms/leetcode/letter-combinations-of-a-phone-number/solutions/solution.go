func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := []string{}
	path := []byte{}

	var dfs func(int)
	dfs = func(i int) {
		if i == len(digits) {
			res = append(res, string(path))
			return
		}

		letters := phone[digits[i]]

		for j := 0; j < len(letters); j++ {
			path = append(path, letters[j])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return res
}