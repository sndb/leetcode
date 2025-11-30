from collections import defaultdict

class Solution:
    def findWinners(self, matches: List[List[int]]) -> List[List[int]]:
        loses = defaultdict(int)
        for w, l in matches:
            if w not in loses:
                loses[w] = 0
            loses[l] += 1

        return [sorted(p for p, l in loses.items() if l == 0),
                sorted(p for p, l in loses.items() if l == 1)]
