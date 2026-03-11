// func findSubstring(s string, words []string) []int {
//     if len(s) == 0 || len(words) == 0 {
//         return []int{}
//     }

//     perms := permuteWords(words)
//     fmt.Println(perms)
//     plen := len(words[0]) * len(words)

//     res := make([]int, 0)
//     for i := 0; i <= len(s)-plen; i++ {
//         for _, p := range perms {
//             found := true
//             for j := 0; j < plen && i+j < len(s); j++ {
//                 if s[i+j] != p[j] {
//                     found = false
//                     break
//                 }
//             }
//             if found {
//                 res = append(res, i)
//                 break
//             }
//         }
//     }

//     return res
// }

// func permuteWords(words []string) []string {
//     res := []string{}
//     n := len(words)

//     used := make([]bool, n)
//     path := []string{}

//     var backtrack func()
//     backtrack = func() {
//         if len(path) == n {
//             s := ""
//             for _, w := range path {
//                 s += w
//             }
//             res = append(res, s)
//             return
//         }

//         for i := 0; i < n; i++ {
//             if used[i] {
//                 continue
//             }

//             used[i] = true
//             path = append(path, words[i])

//             backtrack()

//             path = path[:len(path)-1]
//             used[i] = false
//         }
//     }

//     backtrack()
//     return res
// }

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	res := make([]int, 0)

	wc := make(map[string]int)
	for _, w := range words {
		wc[w]++
	}

	wcount := len(words)
	wlen := len(words[0])
	slen := len(s)

	twlen := wcount * wlen

	for i := 0; i <= slen-twlen; i++ {
		j := 0

		seen := make(map[string]int)

		for j < wcount {
			start := i + (j * wlen)
			ss := s[start : start+wlen]

			if count, ok := wc[ss]; ok {
				seen[ss]++
				if count < seen[ss] {
					break
				}
			} else {
				break
			}
			j++

		}

		if j == wcount {
			res = append(res, i)
		}
	}

	return res
}

