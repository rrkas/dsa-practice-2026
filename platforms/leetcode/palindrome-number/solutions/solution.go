func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	r := 0
	t := x
	for x > 0 {
		d := x % 10
		r = r*10 + d
		x = x / 10
	}
	return r == t
}