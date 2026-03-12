class Solution:
    def trap(self, height):
        l, r = 0, len(height) - 1
        lmax = rmax = 0
        water = 0

        while l < r:
            if height[l] < height[r]:
                if height[l] >= lmax:
                    lmax = height[l]
                else:
                    water += lmax - height[l]
                l += 1
            else:
                if height[r] >= rmax:
                    rmax = height[r]
                else:
                    water += rmax - height[r]
                r -= 1

        return water


class Solution:
    def trap(self, height):
        n = len(height)

        left = [0] * n
        right = [0] * n

        left[0] = height[0]
        for i in range(1, n):
            left[i] = max(left[i - 1], height[i])

        right[n - 1] = height[n - 1]
        for i in range(n - 2, -1, -1):
            right[i] = max(right[i + 1], height[i])

        water = 0
        for i in range(n):
            water += min(left[i], right[i]) - height[i]

        return water


class Solution:
    def trap(self, height):
        stack = []
        water = 0

        for i in range(len(height)):
            while stack and height[i] > height[stack[-1]]:
                bottom = stack.pop()

                if not stack:
                    break

                left = stack[-1]

                width = i - left - 1
                bounded_height = min(height[left], height[i]) - height[bottom]

                water += width * bounded_height

            stack.append(i)

        return water
