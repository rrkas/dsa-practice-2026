func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)

	i := 0

	for i < len(words) {
		line := make([]string, 0)
		wsum := 0

		for i < len(words) && wsum+len(words[i])+len(line) <= maxWidth {
			line = append(line, words[i])
			wsum += len(words[i])
			i++
		}

		spaces := maxWidth - wsum
		gaps := len(line) - 1

		sb := strings.Builder{}

		if i == len(words) || gaps == 0 {
			for wi, w := range line {
				sb.WriteString(w)
				if wi != len(line)-1 {
					sb.WriteString(" ")
				}
			}

			remaining := maxWidth - sb.Len()
			for s := 0; s < remaining; s++ {
				sb.WriteString(" ")
			}
		} else {
			base := spaces / gaps
			extra := spaces % gaps
			for wi, w := range line {
				sb.WriteString(w)
				if wi != len(line)-1 {
					ws := base
					if wi < extra {
						ws++
					}
					for s := 0; s < ws; s++ {
						sb.WriteString(" ")
					}
				}
			}
		}
		res = append(res, sb.String())
	}

	return res
}