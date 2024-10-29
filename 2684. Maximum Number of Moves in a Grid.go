func maxMoves(grid [][]int) int {
    const INF int = int(1e9)

    dp := make([][]int, len(grid))

    for i := 0; i < len(grid); i++ {
        dp[i] = make([]int, len(grid[i]))
    }

    result := 0
    for j := 0; j < len(grid[0]); j++ {
        for i := 0; i < len(grid); i++ {
            if j == 0 {
                dp[i][j] = 0
            } else {
                dp[i][j] = -INF
            }

            // up left
            if i - 1 >= 0 && j - 1 >= 0 && grid[i - 1][j - 1] < grid[i][j] {
                dp[i][j] = max(dp[i][j], 1 + dp[i - 1][j - 1])
            }
            // left
            if j - 1 >= 0 && grid[i][j - 1] < grid[i][j] {
                dp[i][j] = max(dp[i][j], 1 + dp[i][j - 1])
            }

            // down left
            if i + 1 < len(grid) && j - 1 >= 0 && grid[i + 1][j - 1] < grid[i][j] {
                dp[i][j] = max(dp[i][j], 1 + dp[i + 1][j - 1])
            }

            result = max(result, dp[i][j])
        }
    }
    return result
}
