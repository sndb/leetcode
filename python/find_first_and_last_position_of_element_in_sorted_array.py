class Solution:
    def searchRange(self, nums: list[int], target: int) -> list[int]:
        l, r = 0, len(nums) - 1
        while l <= r:
            m = (l + r) // 2
            if nums[m] >= target:
                r = m - 1
            else:
                l = m + 1

        start = l

        l, r = 0, len(nums) - 1
        while l <= r:
            m = (l + r) // 2
            if nums[m] <= target:
                l = m + 1
            else:
                r = m - 1

        end = r

        if end < 0 or nums[end] != target:
            return [-1, -1]

        return [start, end]
