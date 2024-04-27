class Solution:
    def findRotateSteps(self, ring: str, key: str) -> int:
        n = len(ring)
        m = len(key)
        
        ALPHA = 26
        INF = 1000000000

        letterPos = [[] for _ in range(ALPHA)]
        for i in range(n):
            idx = ord(ring[i]) - ord('a')
            letterPos[idx].append(i)

        # dp[next_key_pos][prev_ring_pos]
        dp = [[INF] * n for _ in range(m)]

        idx = ord(key[0]) - ord('a')
        for nextPos in letterPos[idx]:
            delta = nextPos
            delta = min(delta, n - nextPos)
            dp[0][nextPos] = min(dp[0][nextPos], delta + 1)


        for i in range(1, m):
            for j in range(n):

                idx = ord(key[i]) - ord('a')
                for nextPos in letterPos[idx]:
                    delta = abs(nextPos - j)
                    delta = min(delta, n - delta)
                    dp[i][nextPos] = min(dp[i][nextPos], 1 + dp[i - 1][j] + delta)

        result = INF
        for j in range(n):
            result = min(result, dp[m - 1][j])

        return result
 