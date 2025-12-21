class Solution:
    def minDeletionSize(self, strs: list[str]) -> int:
        deletions = set()
        for c in range(len(strs[0])):
            for s1, s2 in zip(strs, strs[1:]):
                if s1[c] > s2[c] and all(s1[i] == s2[i] for i in range(c) if i not in deletions):
                    deletions.add(c)
                    break
        return len(deletions)
