class Solution:
    def largestSquareArea(self, bottomLeft: list[list[int]], topRight: list[list[int]]) -> int:
        result = 0
        for i in range(len(bottomLeft)):
            for j in range(i):
                ax1, ay1 = bottomLeft[i]
                bx1, by1 = topRight[i]
                ax2, ay2 = bottomLeft[j]
                bx2, by2 = topRight[j]
                dx = min(bx1, bx2) - max(ax1, ax2)
                dy = min(by1, by2) - max(ay1, ay2)
                result = max(result, min(dx, dy))
        return result**2
