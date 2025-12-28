class Solution:
    def countNegatives(self, grid: list[list[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        r, c = 0, cols - 1
        result = 0
        while r < rows and c >= 0:
            if grid[r][c] < 0:
                result += rows - r
                c -= 1
            else:
                r += 1
        return result
