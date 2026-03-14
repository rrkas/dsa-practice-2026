class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        n = len(digits)
        carry = 0
        for i in range(n - 1, -1, -1):
            v = digits[i] + carry
            if i == n - 1:
                v += 1
            digits[i] = v % 10
            carry = v // 10

        if carry:
            return [carry, *digits]

        return digits
