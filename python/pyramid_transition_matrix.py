from collections import defaultdict

class Solution:
    def pyramidTransition(self, bottom: str, allowed: list[str]) -> bool:
        trans = defaultdict(list)
        for a, b, c in allowed:
            trans[a + b].append(c)

        def variants(row):
            if len(row) == 2:
                for t in trans[row]:
                    yield t
            else:
                for v in variants(row[1:]):
                    for t in trans[row[:2]]:
                        yield t + v

        def test(row):
            return len(row) == 1 or any(test(v) for v in variants(row))

        return test(bottom)
