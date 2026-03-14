import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	rows := make([]strings.Builder, numRows)

	ridx := 0
	dir := 1

	for i := 0; i < len(s); i++ {
		rows[ridx].WriteByte(s[i])

		if ridx == 0 {
			dir = 1
		} else if ridx == numRows-1 {
			dir = -1
		}

		ridx += dir
	}

	var res strings.Builder
	for i := 0; i < numRows; i++ {
		res.WriteString(rows[i].String())
	}

	return res.String()
}