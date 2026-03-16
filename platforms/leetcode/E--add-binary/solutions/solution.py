class Solution:
    def addBinary(self, a: str, b: str) -> str:
        res = ""
        carry = 0
        m, n = len(a), len(b)
        i, j = m - 1, n - 1
        while i >= 0 and j >= 0:
            s = ord(a[i]) + ord(b[j]) - 2 * ord("0") + carry
            carry = s // 2
            s = s % 2
            res = str(s) + res
            i -= 1
            j -= 1

        while i >= 0:
            s = ord(a[i]) - ord("0") + carry
            carry = s // 2
            s = s % 2
            res = str(s) + res
            i -= 1

        while j >= 0:
            s = ord(b[j]) - ord("0") + carry
            carry = s // 2
            s = s % 2
            res = str(s) + res
            j -= 1

        if carry:
            res = str(carry) + res

        return res
