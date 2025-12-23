from functools import cache

class Solution:
    def minDeletionSize(self, strs: list[str]) -> int:
        @cache
        def solve(c, prev):
            if c == len(strs[0]):
                return 0

            if prev is not None and any(strs[r][prev] > strs[r][c] for r in range(len(strs))):
                return 1 + solve(c + 1, prev)

            return min(solve(c + 1, c), 1 + solve(c + 1, prev))

        return solve(0, None)
