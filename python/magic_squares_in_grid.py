class Solution:
    def numMagicSquaresInside(self, grid: list[list[int]]) -> int:
        result = 0
        for r in range(len(grid) - 2):
            for c in range(len(grid[0]) - 2):
                nums = {grid[r + i][c + j] for i in range(3) for j in range(3)}
                result += len(nums) == 9 and all(1 <= x <= 9 for x in nums) and len({
                    sum(grid[r + i][c + i] for i in range(3)),
                    sum(grid[r + i][c + 2 - i] for i in range(3)),
                    *(sum(grid[r + j][c + i] for i in range(3)) for j in range(3)),
                    *(sum(grid[r + i][c + j] for i in range(3)) for j in range(3)),
                }) == 1
        return result
