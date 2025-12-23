class Solution:
    def maxTwoEvents(self, events: list[list[int]]) -> int:
        end_sorted = sorted(events, key=lambda x: x[1])
        i, prev_max, result = 0, 0, 0
        for start, _, value in sorted(events):
            while end_sorted[i][1] < start:
                prev_max = max(prev_max, end_sorted[i][2])
                i += 1

            result = max(result, prev_max + value)

        return result
