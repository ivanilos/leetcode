class Solution:
    def removeKdigits(self, num: str, k: int) -> str:
        result = ""

        for digit in num:
            while len(result) > 0 and result[-1] > digit and k > 0:
                k -= 1
                result = result[:-1]
            
            result += digit

        while k > 0:
            result = result[:-1]
            k -= 1

        if len(result) == 0:
            return "0"
        else:
            start = 0
            while start < len(result) and result[start] == '0':
                start += 1

            if start == len(result):
                return "0"
            else:
                result = result[start::]
                return result
