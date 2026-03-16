import "strconv"
import "strings"

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}

	rle := func(is string) string {
		t := strings.Builder{}
		i := 0
		for i < len(is) {
			char := is[i]
			count := 1
			for i+1 < len(is) && is[i] == is[i+1] {
				i++
				count++
			}
			t.WriteString(strconv.Itoa(count) + string(char))
			i++
		}
		return t.String()
	}

	s := "1"
	for i := 1; i < n; i++ {
		s = rle(s)
	}

	return s
}