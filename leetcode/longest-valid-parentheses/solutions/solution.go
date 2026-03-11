// func longestValidParentheses(s string) int {
// 	stack := []int{-1}
// 	maxLen := 0

// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '(' {
// 			stack = append(stack, i)
// 		} else {
// 			stack = stack[:len(stack)-1]

// 			if len(stack) == 0 {
// 				stack = append(stack, i)
// 			} else {
// 				length := i - stack[len(stack)-1]
// 				if length > maxLen {
// 					maxLen = length
// 				}
// 			}
// 		}
// 	}
// 	return maxLen
// }

///////////////////////////////////////////////////////////////////////////////////////

// func longestValidParentheses(s string) int {
//     left, right, maxLen := 0, 0, 0

//     for i := 0; i < len(s); i++ {
//         if s[i] == '(' {
//             left++
//         } else {
//             right++
//         }

//         if left == right {
//             if 2*right > maxLen {
//                 maxLen = 2 * right
//             }
//         } else if right > left {
//             left, right = 0, 0
//         }
//     }

//     left, right = 0, 0

//     for i := len(s)-1; i >= 0; i-- {
//         if s[i] == '(' {
//             left++
//         } else {
//             right++
//         }

//         if left == right {
//             if 2*left > maxLen {
//                 maxLen = 2 * left
//             }
//         } else if left > right {
//             left, right = 0, 0
//         }
//     }

//     return maxLen
// }

///////////////////////////////////////////////////////////////////////////

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	maxLen := 0

	for i := 1; i < n; i++ {
		if s[i] == ')' {

			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}

			} else {
				j := i - dp[i-1] - 1
				if j >= 0 && s[j] == '(' {
					dp[i] = dp[i-1] + 2
					if j-1 >= 0 {
						dp[i] += dp[j-1]
					}
				}
			}

			if dp[i] > maxLen {
				maxLen = dp[i]
			}
		}
	}

	return maxLen
}

