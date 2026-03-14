import "strings"

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	i, j := m-1, n-1
	var carry byte = 0
	resbin := make([]byte, max(m, n)+1)
	reslen := 0

	for i >= 0 && j >= 0 {
		v := (a[i] - '0') + (b[j] - '0') + carry
		carry = v / 2
		v = v % 2
		resbin[reslen] = v

		i--
		j--
		reslen++
	}

	for i >= 0 {
		v := (a[i] - '0') + carry
		carry = v / 2
		v = v % 2
		resbin[reslen] = v

		i--
		reslen++
	}

	for j >= 0 {
		v := (b[j] - '0') + carry
		carry = v / 2
		v = v % 2
		resbin[reslen] = v

		i--
		reslen++
	}

	if carry > 0 {
		resbin[reslen] = carry
		reslen++
	}

	res := strings.Builder{}
	reslen--
	for reslen >= 0 {
		res.WriteByte('0' + resbin[reslen])
		reslen--
	}
	return res.String()
}