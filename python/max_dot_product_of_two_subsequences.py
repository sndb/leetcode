from functools import cache

class Solution:
    def maxDotProduct(self, nums1: list[int], nums2: list[int]) -> int:
        @cache
        def solve(i1, i2):
            if i1 >= len(nums1) or i2 >= len(nums2):
                return 0
            return max(solve(i1 + 1, i2 + 1) + nums1[i1] * nums2[i2],
                       solve(i1 + 1, i2),
                       solve(i1, i2 + 1))

        result = solve(0, 0)
        if result == 0:
            return max(max(nums1) * max(nums2),
                       max(nums1) * min(nums2),
                       min(nums1) * max(nums2),
                       min(nums1) * min(nums2))
        return result
