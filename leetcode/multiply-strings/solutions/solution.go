import "strings"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	l1, l2 := len(num1), len(num2)
	prod := make([]int, l1+l2)

	for i1 := l1 - 1; i1 >= 0; i1-- {
		for i2 := l2 - 1; i2 >= 0; i2-- {
			curr := i1 + i2 + 1
			carry := i1 + i2

			tot := int(num1[i1]-'0')*int(num2[i2]-'0') + prod[curr]
			prod[curr] = tot % 10
			prod[carry] += tot / 10
		}
	}

	s := strings.Builder{}
	lzero := true
	for _, d := range prod {
		if d == 0 && lzero {
			continue
		}
		lzero = false
		s.WriteByte('0' + byte(d))
	}

	if s.Len() == 0 {
		return "0"
	}

	return s.String()
}