class Solution:
    def maximalRectangle(self, matrix: list[list[str]]) -> int:
        rows, cols = len(matrix), len(matrix[0])
        rectangles = [[[] for _ in range(cols)] for _ in range(rows + 1)]
        for r in range(rows - 1, -1, -1):
            run = 0
            for c in range(cols - 1, -1, -1):
                if matrix[r][c] == "1":
                    run += 1
                else:
                    run = 0
                    continue

                rectangles[r][c].append((1, run))
                for low_rows, low_cols in rectangles[r + 1][c]:
                    rectangles[r][c].append((low_rows + 1, min(run, low_cols)))

        result = 0
        for row in rectangles:
            for sizes in row:
                for a, b in sizes:
                    result = max(result, a * b)
        return result
