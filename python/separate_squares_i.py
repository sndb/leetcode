class Solution:
    def separateSquares(self, squares: list[list[int]]) -> float:
        low = min(y for _, y, _ in squares)
        high = max(y + l for _, y, l in squares)

        while abs(low - high) > 10**-5:
            below, above = 0, 0
            mid = (low + high) / 2
            for _, y, l in squares:
                if y >= mid:
                    above += l * l
                elif y + l <= mid:
                    below += l * l
                else:
                    below += (half := l * (mid - y))
                    above += l * l - half

            if above > below:
                low = mid
            else:
                high = mid

        return high
