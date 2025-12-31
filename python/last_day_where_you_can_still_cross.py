class Solution:
    def latestDayToCross(self, row: int, col: int, cells: list[list[int]]) -> int:
        lo, hi = 0, len(cells) - 1
        while lo <= hi:
            mi = (lo + hi) // 2
            water = set(map(tuple, cells[: mi + 1]))
            queue = [(1, c) for c in range(1, col + 1) if (1, c) not in water]
            seen = set(queue)
            while queue:
                r, c = queue.pop()
                if r == row:
                    lo = mi + 1
                    break

                for nr, nc in [(r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)]:
                    if 1 <= nr <= row and 1 <= nc <= col and (nr, nc) not in water and (nr, nc) not in seen:
                        queue.append((nr, nc))
                        seen.add((nr, nc))
            else:
                hi = mi - 1

        return lo
