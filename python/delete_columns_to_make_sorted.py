class Solution:
    def minDeletionSize(self, strs: list[str]) -> int:
        return sum(any(a > b for a, b in zip(s, s[1:])) for s in zip(*strs))
