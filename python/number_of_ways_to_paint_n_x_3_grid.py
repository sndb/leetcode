from itertools import product

class Solution:
    def numOfWays(self, n: int) -> int:
        combinations = [x for x in product([1, 2, 3], repeat=3) if x[0] != x[1] != x[2]]

        continuations = {
            a: [b for b in combinations if not any(x == y for x, y in zip(a, b))]
            for a in combinations
        }

        counts = {x: 1 for x in combinations}
        for _ in range(n - 1):
            next = {x: 0 for x in combinations}
            for a, n in counts.items():
                for b in continuations[a]:
                    next[b] += n
            counts = next

        return sum(counts.values()) % (10**9 + 7)
