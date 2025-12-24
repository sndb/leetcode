from itertools import accumulate

class Solution:
    def minimumBoxes(self, apple: list[int], capacity: list[int]) -> int:
        return next(i + 1 for i, c in enumerate(accumulate(sorted(capacity, reverse=True))) if c >= sum(apple))
