from functools import cache

class Solution:
    def minimumDeleteSum(self, s1: str, s2: str) -> int:
        @cache
        def solve(i1, i2):
            if i1 >= len(s1):
                return sum(map(ord, s2[i2:]))
            if i2 >= len(s2):
                return sum(map(ord, s1[i1:]))
            if s1[i1] == s2[i2]:
                return solve(i1 + 1, i2 + 1)
            return min(
                ord(s1[i1]) + solve(i1 + 1, i2),
                ord(s2[i2]) + solve(i1, i2 + 1),
            )

        return solve(0, 0)
