class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        # if "0" in [num1, num2]:
        #     return "0"

        # m, n = len(num1), len(num2)
        # prod = [0] * (m + n)
        # for i in range(m - 1, -1, -1):
        #     for j in range(n - 1, -1, -1):
        #         curr = i + j + 1
        #         carry = i + j
        #         tot = int(num1[i]) * int(num2[j]) + prod[curr]
        #         prod[curr] = tot % 10
        #         prod[carry] += tot // 10

        # result = "".join(map(str, prod)).lstrip("0")
        # return result if result else "0"

        return str(int(num1) * int(num2))
