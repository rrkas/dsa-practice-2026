class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if len(s) == 1 or numRows == 1:
            return s

        rows = [""] * numRows

        ridx = 0
        dirn = 1  # 1 = down, -1 = up

        for c in s:
            rows[ridx] += c
            ridx += dirn
            if ridx == numRows - 1 or ridx == 0:
                dirn *= -1

        return "".join(rows)
