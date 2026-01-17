class Solution:
    def maximizeSquareArea(self, m: int, n: int, hFences: list[int], vFences: list[int]) -> int:
        hfs = [1] + sorted(hFences) + [m]
        vfs = [1] + sorted(vFences) + [n]

        sides = set()
        for i in range(len(hfs)):
            for j in range(i):
                sides.add(hfs[i] - hfs[j])

        result = 0
        for i in range(len(vfs)):
            for j in range(i):
                if (side := vfs[i] - vfs[j]) in sides:
                    result = max(result, side)

        return (result**2) % (10**9 + 7) if result != 0 else -1
