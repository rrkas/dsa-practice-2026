def maximalRectangleArea(heights) -> int:
    heights.append(0)
    stack = []
    maxa = 0

    for i in range(len(heights)):
        while stack and heights[stack[-1]] > heights[i]:
            h = heights[stack.pop()]
            width = i - stack[-1] - 1 if stack else i
            maxa = max(maxa, h * width)
        stack.append(i)

    return maxa


class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        m, n = len(matrix), len(matrix[0])
        heights = [0] * n
        ans = 0

        for i in range(m):
            for j in range(n):
                if matrix[i][j] == "1":
                    heights[j] += 1
                else:
                    heights[j] = 0

            ans = max(ans, maximalRectangleArea(heights))

        return ans
