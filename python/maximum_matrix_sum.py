class Solution:
    def maxMatrixSum(self, matrix: list[list[int]]) -> int:
        sum, neg, low = 0, 0, 1e9
        for row in matrix:
            for x in row:
                if x < 0:
                    neg += 1
                    x = -x
                sum += x
                low = min(low, x)
        if neg % 2 == 0:
            return sum
        return sum - 2 * low
