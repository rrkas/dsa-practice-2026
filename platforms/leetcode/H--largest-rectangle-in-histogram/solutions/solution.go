func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	maxa := 0
	n := len(heights)

	for i := 0; i <= n; i++ {
		curr_h := 0
		if i < n {
			curr_h = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] > curr_h {
			popped_idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[popped_idx]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			maxa = max(maxa, h*width)
		}

		stack = append(stack, i)
	}

	return maxa
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