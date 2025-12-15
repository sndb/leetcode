class Solution:
    def getDescentPeriods(self, prices: List[int]) -> int:
        prev = 0
        count = 0
        result = 0
        for p in prices:
            if p == prev - 1:
                count += 1
            else:
                count = 1
            result += count
            prev = p
        return result
