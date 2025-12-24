class Solution:
    def generateParenthesis(self, n: int) -> list[str]:
        def f(open, close):
            if open == 0:
                yield ")" * close
                return

            if close > 0:
                for r in f(open, close - 1):
                    yield ")" + r

            for r in f(open - 1, close + 1):
                yield "(" + r

        return list(f(n, 0))
