class Solution:
    def maximizeSquareHoleArea(self, n: int, m: int, hBars: list[int], vBars: list[int]) -> int:
        hBars.sort()
        vBars.sort()

        hmax, hcur = 0, 0
        for i in range(1, len(hBars)):
            if hBars[i] - hBars[i - 1] == 1:
                hcur += 1
                hmax = max(hmax, hcur)
            else:
                hcur = 0

        vmax, vcur = 0, 0
        for i in range(1, len(vBars)):
            if vBars[i] - vBars[i - 1] == 1:
                vcur += 1
                vmax = max(vmax, vcur)
            else:
                vcur = 0

        return min(vmax + 2, hmax + 2) ** 2
