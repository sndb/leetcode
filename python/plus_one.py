class Solution:
    def plusOne(self, digits: list[int]) -> list[int]:
        for i in reversed(range(len(digits))):
            if digits[i] < 9:
                digits[i] += 1
                return digits

            digits[i] = 0

        return [1] + digits
