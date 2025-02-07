func minPathSum(grid [][]int) int {
    dp := make([][]int, len(grid))
    for i := 0; i < len(grid); i++ {
        dp[i] = make([]int, len(grid[i]))
    }

    dp[0][0] = grid[0][0]
    for i := 1; i < len(grid); i++ {
        dp[i][0] = grid[i][0] + dp[i - 1][0]
    }
    for j := 1; j < len(grid[0]); j++ {
        dp[0][j] = grid[0][j] + dp[0][j - 1]
    }

    for i := 1; i < len(grid); i++ {
        for j := 1; j < len(grid[i]); j++ {
            dp[i][j] = grid[i][j] + min(dp[i - 1][j], dp[i][j - 1])
        }
    }

    return dp[len(grid) - 1][len(grid[0]) - 1]
}
