class Solution:
    # @param A : tuple of integers
    # @param B : tuple of integers
    # @return an integer
    def canCompleteCircuit(self, A, B):
        total = 0
        tank = 0
        start = 0

        for i in range(len(A)):
            diff = A[i] - B[i]
            total += diff
            tank += diff

            # cannot reach next station
            if tank < 0:
                start = i + 1
                tank = 0

        return start if total >= 0 else -1
