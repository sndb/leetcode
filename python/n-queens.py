class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        result = []
        queens = []
        cols = [False] * n
        dia1 = [False] * (2 * n)
        dia2 = [False] * (2 * n)

        def solve(r):
            if r == n:
                board = [["." for _ in range(n)] for _ in range(n)]
                for qr, qc in queens:
                    board[qr][qc] = "Q"

                result.append(["".join(row) for row in board])
                return

            for c in range(n):
                d1, d2 = r - c, n - r - c
                if not cols[c] and not dia1[d1] and not dia2[d2]:
                    cols[c], dia1[d1], dia2[d2] = True, True, True
                    queens.append((r, c))
                    solve(r + 1)

                    cols[c], dia1[d1], dia2[d2] = False, False, False
                    queens.pop()

        solve(0)
        return result
