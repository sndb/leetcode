from math import isqrt

class Solution:
    def sumFourDivisors(self, nums: list[int]) -> int:
        result = 0
        for n in nums:
            divs = {1, n}
            for d in range(2, isqrt(n) + 1):
                if n % d == 0:
                    divs |= {d, n // d}
                if len(divs) > 4:
                    break
            if len(divs) == 4:
                result += sum(divs)
        return result
