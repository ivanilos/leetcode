class Solution:
    def longestIdealString(self, s: str, k: int) -> int:
        dp = [[0] * 26 for _ in range(len(s))]

        firstLetterIdx = ord(s[0]) - ord('a')
        dp[0][firstLetterIdx] = 1

        i = 1
        while i < len(s):
            idx = ord(s[i]) - ord('a')
            small = max(0, idx - k)
            big = min(25, idx + k)

            j = 0
            while j < 26:
                dp[i][j] = dp[i - 1][j]
                j += 1

            j = small
            while j <= big:
                dp[i][idx] = max(dp[i][idx], 1 + dp[i - 1][j])
                j += 1

            i += 1

        result = 0
        j = 0
        while j < 26:
            result = max(result, dp[len(s) - 1][j])
            j += 1

        return result
