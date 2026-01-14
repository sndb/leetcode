class Solution:
    def merge(self, intervals: list[list[int]]) -> list[list[int]]:
        intervals.sort()
        result = [intervals[0]]
        for start, end in intervals[1:]:
            last = result[-1]
            if start <= last[1]:
                last[1] = max(last[1], end)
            else:
                result.append([start, end])
        return result
