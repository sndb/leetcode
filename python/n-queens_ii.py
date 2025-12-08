class Solution:
    def totalNQueens(self, n: int) -> int:
        queens = []
        cols = [False] * n
        dia1 = [False] * (2 * n)
        dia2 = [False] * (2 * n)

        def solve(r):
            if r == n:
                return 1

            result = 0
            for c in range(n):
                d1, d2 = r - c, n - r - c
                if not cols[c] and not dia1[d1] and not dia2[d2]:
                    cols[c], dia1[d1], dia2[d2] = True, True, True
                    queens.append((r, c))
                    result += solve(r + 1)

                    cols[c], dia1[d1], dia2[d2] = False, False, False
                    queens.pop()

            return result

        return solve(0)
