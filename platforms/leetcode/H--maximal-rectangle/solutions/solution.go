func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)
	maxArea := 0

	for i := 0; i < m; i++ {
		// build histogram
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		// compute largest rectangle in histogram
		maxArea = max(maxArea, largestRectangleArea(heights))
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0) // sentinel
	stack := make([]int, len(heights)+1)
	stackn := 0
	maxa := 0

	for i := 0; i < len(heights); i++ {
		for stackn > 0 && heights[stack[stackn-1]] > heights[i] {
			top := stack[stackn-1]
			stackn--

			h := heights[top]

			width := i
			if stackn > 0 {
				width = i - stack[stackn-1] - 1
			}

			area := h * width
			if area > maxa {
				maxa = area
			}
		}
		stack[stackn] = i
		stackn++
	}

	return maxa
}