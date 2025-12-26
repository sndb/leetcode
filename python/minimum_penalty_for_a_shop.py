class Solution:
    def bestClosingTime(self, customers: str) -> int:
        idx, acc, best = 0, 0, 0
        for i, c in enumerate(customers):
            acc += c == "Y" or -1
            if acc > best:
                idx = i + 1
                best = acc
        return idx
