class Solution:
    def minFallingPathSum(self, grid: List[List[int]]) -> int:
        n = len(grid)
        INF = 100000000

        if n == 1:
            return grid[0][0]

        dp = [[0] * n for _ in range(n)]

        for j in range(0, n):
            dp[0][j] = grid[0][j]

        for i in range(1, n):
            mini = [(INF, 0), (INF, 1)]
            for j in range(0, n):
                if dp[i - 1][j] < mini[0][0]:
                    mini[1] = mini[0]
                    mini[0] = (dp[i - 1][j], j)
                elif dp[i - 1][j] < mini[1][0]:
                    mini[1] = (dp[i - 1][j], j)

            for j in range(0, n):
                dp[i][j] = grid[i][j] + mini[0][0] if mini[0][1] != j else grid[i][j] + mini[1][0]

        result = INF
        for j in range(0, n):
            result = min(result, dp[n - 1][j])

        return result
