class Solution:
    def maximumOddBinaryNumber(self, s: str) -> str:
        ones = 0
        zeroes = 0
        for ch in s:
            if ch == '0':
                zeroes += 1
            else:
                ones += 1

        result = '1' * (ones - 1) + '0' * zeroes + '1'

        return result 
